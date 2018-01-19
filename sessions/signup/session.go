package signup

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session-l")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session-l",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}

	return u
}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session-l")

	if err != nil {
		return false
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	s := dbSessions[c.Value]
	_, ok := dbUsers[s.un]
	return ok
}

func cleanSessions() {
	fmt.Println("Pre Clean")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second*30) || v.un == "" {
			delete(dbSessions, k)
		}
	}

	dbSessionsCleaned = time.Now()
	fmt.Println("AFter Clean")
	showSessions()
}

func showSessions() {
	fmt.Println("********* showing sessions ***********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
