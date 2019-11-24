package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	lau := person{
		firstName: "Laura",
		lastName:  "Purvis",
		contact: contactInfo{
			email:   "lau@mail.com",
			zipCode: 12345,
		},
	}

	lau.updateName("lauraaaaaa")
	lau.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
