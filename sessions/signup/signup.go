package signup

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("./signup/templates/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["test@test.com"] = user{"test@test.com", bs, "Zee", "Bee"}
	dbSessionsCleaned = time.Now()
}

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/bar", bar)
	r.HandleFunc("/signup", isLoggedIn(getSignup)).Methods("GET")
	r.HandleFunc("/signup", isLoggedIn(postSignup)).Methods("POST")
	r.HandleFunc("/login", isLoggedIn(getLogin)).Methods("GET")
	r.HandleFunc("/login", isLoggedIn(postLogin)).Methods("POST")
	r.HandleFunc("/logout", isLoggedIn(getLogout)).Methods("GET")
	r.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", r)
	fmt.Println("Now serving on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func isLoggedIn(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if alreadyLoggedIn(w, r) {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		h.ServeHTTP(w, r)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	showSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	un := r.FormValue("username")
	p := r.FormValue("password")
	u, ok := dbUsers[un]
	if !ok {
		http.Error(w, "Username and/or password to no match", http.StatusForbidden)
		return
	}

	err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
	if err != nil {
		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		return
	}

	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session-l",
		Value: sID.String(),
	}
	http.SetCookie(w, c)
	dbSessions[c.Value] = session{un, time.Now()}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func getLogout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session-l")
	delete(dbSessions, c.Value)

	c = &http.Cookie{
		Name:   "session-l",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func postSignup(w http.ResponseWriter, r *http.Request) {
	un := r.FormValue("username")
	p := r.FormValue("password")
	f := r.FormValue("firstname")
	l := r.FormValue("lastname")

	if _, ok := dbUsers[un]; ok {
		http.Error(w, "Username already taken", http.StatusForbidden)
	}

	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session-l",
		Value: sID.String(),
	}
	http.SetCookie(w, c)
	dbSessions[c.Value] = session{un, time.Now()}
	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	u := user{un, bs, f, l}
	dbUsers[un] = u

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func getSignup(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
