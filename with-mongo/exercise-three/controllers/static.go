package controllers

import (
	"net/http"
	"html/template"

	"../session"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (c Controller) Bar(w http.ResponseWriter, r *http.Request) {
	u : session.GetUser(w, r)
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r "/", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
	}

	session.Show()
	c.tpl.ExecuteTemplate("w, "bar.gohtml", u)
}