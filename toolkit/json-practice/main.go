package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	FirstName string
	LastName  string
	Items     []string
}

type thumbnail struct {
	URL           string
	Height, Width int
}

type img struct {
	Height, Width int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

type city struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Name      string  `json:"City"`
	State     string  `json:"State"`
	Country   string  `json:"Country"`
	Zip       string  `json:"Zip"`
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marshall", marshall)
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/unmarshall", unmarshall)
	http.HandleFunc("/cities", cities)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>FOO</title>
			</head>
			<body>
			You are at foo
			</body>
			</html>`
	w.Write([]byte(s))
}

func marshall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		FirstName: "Jim",
		LastName:  "Bond",
		Items:     []string{"martini", "car"},
	}

	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}

func unmarshall(w http.ResponseWriter, r *http.Request) {
	var data img
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}
	fmt.Println(data)
	for i, v := range data.IDs {
		fmt.Println(i, v)
	}
	fmt.Println(data.Thumbnail.URL)
}

func cities(w http.ResponseWriter, r *http.Request) {
	var cities []city
	rcvd := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	err := json.Unmarshal([]byte(rcvd), &cities)
	if err != nil {
		log.Fatalln(err)
	}

	for _, c := range cities {
		fmt.Println("Name: ", c.Name, "State: ", c.State)
	}

}

func encode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		FirstName: "Zac",
		LastName:  "Baston",
		Items:     []string{"macbook", "barbell"},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
