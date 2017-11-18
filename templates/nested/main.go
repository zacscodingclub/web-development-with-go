package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

type weightlifter struct {
	person
	EvenLifts bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	p := person{
		Name: "Zac",
		Age:  33,
	}

	_ = p
	w := weightlifter{
		person{
			Name: "Zac",
			Age:  33,
		},
		true,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", w)
	if err != nil {
		log.Fatalln(err)
	}
}
