package main

import "fmt"

func main() {
	compositeLiteral()
	p := person{name: "Zac", age: 33, address: "Home"}
	p.speak()
}

func compositeLiteral() {
	xi := []int{2, 4, 6, 7, 9}
	fmt.Println(xi)

	m := map[string]int{
		"Zac":     33,
		"Roxanne": 30,
	}

	for k, v := range m {
		fmt.Println(k, "\t", v)
	}
}

type person struct {
	name    string
	age     int
	address string
}

func (p person) speak() {
	fmt.Println("Hi, I'm ", p.name, "from", p.address)
}
