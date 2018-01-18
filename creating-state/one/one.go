package one

import (
	"fmt"
	"io"
	"net/http"
)

func Run() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("Now serving at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q") // FormValue retrieves query params
	if v == "" {
		io.WriteString(w, "No search params entered")
		return
	}
	io.WriteString(w, "Do my search: "+v)
}
