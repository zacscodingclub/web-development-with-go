package main

import (
	"fmt"
	"net/http"

	"./controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Index!</title>
		</head>
		<body>
			<a href="/user/1">User One</a>	
		</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}
