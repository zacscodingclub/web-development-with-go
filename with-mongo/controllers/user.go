package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/julienschmidt/httprouter"
	"github.com/zacscodingclub/web-development-with-go/with-mongo/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "111"
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleting the user")
}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Zac",
		Gender: "Male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	fmt.Println(p)
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}
