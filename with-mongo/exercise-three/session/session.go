package session

import (
	"fmt"
	"net/http"
	"time"

	"../models"

	uuid "github.com/satori/go.uuid"
)

const Length int = 30

var Users = map[string]models.User{}
var Sessions = map[string]models.Session{}
var LastCleaned time.Time = time.Now()

func GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = Length
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if s, ok := Sessions[c.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
		u = Users[s.UserName]
	}
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
	}
	_, ok = Users[s.UserName]
	// refresh session
	c.MaxAge = Length
	http.SetCookie(w, c)
	return ok
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	Show()                      // for demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	Show()                     // for demonstration purposes
}

// for demonstration purposes
func Show() {
	fmt.Println("********")
	for k, v := range Sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
