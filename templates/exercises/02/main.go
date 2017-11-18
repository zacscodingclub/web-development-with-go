package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name, Address, City, Zipcode, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))

}

func main() {
	h := hotels{
		hotel{
			Name:    "Motif Seattle",
			Address: "1415 Fifth Avenue",
			City:    "Seattle",
			Zipcode: "98101-2313",
			Region:  "Downtown",
		},
		hotel{
			Name: "Inn at the Market", Address: "86 Pine St.", City: "Seattle", Zipcode: "98101-1571", Region: "Pioneer Square",
		},
		hotel{
			Name: "Four Seasons Hotel Seattle", Address: "99 Union Street", City: "Seattle", Zipcode: "98101-5011", Region: "Downtown",
		},
		hotel{
			Name: "The Paramount Hotel", Address: "724 Pine St.", City: "Seattle", Zipcode: "98101-1843", Region: "First Hill",
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", h)
	if err != nil {
		log.Fatalln(err)
	}
}
