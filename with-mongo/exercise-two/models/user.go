package models

import (
	"encoding/json"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func StoreUsers(m map[string]User) {
	f, err := os.Create("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(m)
}

func LoadUsers() map[string] {
	m := make(map[string]User)

	f, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return m
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}

	return m
}