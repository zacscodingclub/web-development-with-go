package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	hmacRun()
	base64Run()
	basicContext()
}

func basicContext() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/more", more)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

func more(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "user", 1)
	ctx = context.WithValue(ctx, "fname", "Zac")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		uid := ctx.Value("user").(int)
		time.Sleep(4 * time.Second)
		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}

}

func base64Run() {
	s := "Bacon ipsum dolor amet salami chicken pork, leberkas corned beef ham sirloin shank meatball prosciutto turkey chuck. Kevin brisket chicken picanha, ham pancetta strip steak landjaeger short loin pork belly pork loin. Landjaeger porchetta pork corned beef, pork chop beef ribeye boudin short ribs short loin tri-tip shank. Shankle sirloin biltong ball tip turkey pork loin tenderloin hamburger flank ground round buffalo pork."
	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("decode broke")
	}
	fmt.Println(string(bs))
}

func hmacRun() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@example.co")
	fmt.Println(c)
}

func getCode(s string) string {
	key := []byte("super-secret-key")
	h := hmac.New(sha256.New, key)
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
