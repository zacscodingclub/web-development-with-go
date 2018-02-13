package session

import (
	"fmt"
	"net/http"
	"time"

	"../models"
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
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if s, ok := Sessions[c.Value]; ok {
		s.lastActivity = time.Now()
		Sessions[c.Value] = s
		u = Users[s.un]
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
		s.lastActivity = time.Now()
		Sessions[c.Value] = s
	}
	_, ok = Users[s.un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	Show()                      // for demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
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
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
