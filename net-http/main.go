package main

import (
	"net/http"
)

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

func main() {

	http.ListenAndServe(":3000", "Hello World")
}
