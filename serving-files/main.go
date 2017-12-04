package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	static()
}

func static() {
	http.ListenAndServe(":3000", http.FileServer(http.Dir("./public")))
}

func serveDir() {
	dog := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, `
			<img src="/dog.jpg" />
		`)
	}

	dogFile := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/447H.jpg")
	}

	http.Handle("/", http.FileServer(http.Dir("../..")))
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", dogFile)
	http.ListenAndServe(":3000", nil)
}
func localDog() {
	dog := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, `
			<img src="/dog.jpg" />
		`)
	}

	dogPic := func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public/447H.jpg")
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			http.Error(w, "file not found", 404)
			return
		}

		http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
	}

	dogFile := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/447H.jpg")
	}
	_ = dogPic

	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", dogFile)
	http.ListenAndServe(":3000", nil)
}

func notServing() {
	me := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, `
			<!-- not serving from this server -->
			<img src="https://cdn.gratisography.com/photos/447H.jpg" />
			<img src="/447H.jpg" />
		`)
	}

	http.HandleFunc("/", me)
	http.ListenAndServe(":3000", nil)
}
