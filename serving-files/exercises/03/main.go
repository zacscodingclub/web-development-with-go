package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public/images"))
	http.Handle("/pics/", http.StripPrefix("/pics", fs))
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, error := template.ParseFiles("./templates/index.gohtml")
	if error != nil {
		log.Fatalln(error)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
