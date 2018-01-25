package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("Now serving at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("photo")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		path := filepath.Join(wd, "public", "photos")

		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		}

		path = filepath.Join(wd, "public", "photos", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		mf.Seek(0, 0)
		io.Copy(nf, mf)
		c = appendValue(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value

	if !strings.Contains(s, fname) {
		s += "|" + fname
	}

	c.Value = s

	http.SetCookie(w, c)
	return c
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("ph-session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "ph-session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
