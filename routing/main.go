package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat", cat)
	http.ListenAndServe(":3000", nil)
}

func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "purr, meow, meow")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "woof, woof, woof")
}
