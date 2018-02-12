package main

import (
	"fmt"
	"net/http"

	"./controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Println("Now serving on localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}
