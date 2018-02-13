package main

import (
	"fmt"
	"html/template"
	"net/http"

	"./controllers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	c := controllers.NewController(tpl)
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/signup", c.Signup)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("Now serving at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
