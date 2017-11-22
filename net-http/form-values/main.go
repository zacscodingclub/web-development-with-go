package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type page int

type data struct {
	Method        string
	URL           *url.URL
	Submissions   map[string][]string
	Header        http.Header
	ContentLength int64
}

func (p page) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	d := data{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", d)
}
func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}
func main() {
	var p page
	http.ListenAndServe(":3000", p)
}
