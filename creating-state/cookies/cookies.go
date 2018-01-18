package cookies

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("Now serving on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func 

func set(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "my-cookie",
		Value: "something unique",
	}
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Check browser for cookie")
	fmt.Fprintln(w, "In chrome: dev tools -> application -> cookies")
}

func abundance(w http.ResponseWriter, r *http.Request) {
	c2 := &http.Cookie{
		Name:  "my-cookie1",
		Value: "something unique1",
	}
	c3 := &http.Cookie{
		Name:  "my-cookie2",
		Value: "something unique2",
	}
	http.SetCookie(w, c2)
	http.SetCookie(w, c3)
	fmt.Fprintln(w, "Check browser for cookie")
	fmt.Fprintln(w, "In chrome: dev tools -> application -> cookies")
}

func read(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Cookie #1:", c1)
	}

	c2, err := r.Cookie("my-cookie1")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Cookie #2:", c2)
	}

	c3, err := r.Cookie("my-cookie2")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Cookie #3:", c3)
	}

}
