package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogPhoto)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<h1>Foo Ran</h1>
	`)
}

func dogPhoto(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../public/447H.jpg")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl, error := template.ParseFiles("dog.gohtml")
	if error != nil {
		log.Fatalln(error)
	}

	tpl.ExecuteTemplate(w, "dog.gohtml", "Roofus")
}
