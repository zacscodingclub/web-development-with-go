package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	port := ":3000"
	http.ListenAndServe(port, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the index")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Woof woof bro")
}
func me(w http.ResponseWriter, r *http.Request) {
	templateName := "index.gohtml"
	tpl, err := template.ParseFiles(templateName)
	if err != nil {
		log.Fatalln("error parsing template:", err)
	}
	data := "Zac"

	err = tpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Fatalln("error executing template:", err)
	}
}
