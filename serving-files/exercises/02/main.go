package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/pic/surf.jpg", surfPhoto)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func surfPhoto(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/images/surf.jpg")
}
