package main

import (
	"io"
	"net/http"
)

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
	io.WriteString(w, "My name is Zac")
}
