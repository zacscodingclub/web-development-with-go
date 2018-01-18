package two

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Run() {
	http.HandleFunc("/", root)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("Now serving at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	first := r.FormValue("first_name")
	last := r.FormValue("last_name")
	subbed := r.FormValue("subscribe") == "on"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
		<label for="first_name">First Name</label>
		<input type="text" name="first_name">
		<label for="last_name">Last Name</label>
		<input type="text" name="last_name">
		<label for="subscribe">Subscribe</label>
		<input type="checkbox" name="subscribe">
		<input type="submit">
	</form>
	<br>
	<h2>First: `+first+`</h2>
	<h2>Last: `+last+`</h2>
	<h2>Subbed? `+strconv.FormatBool(subbed)+`</h2>`)
}
