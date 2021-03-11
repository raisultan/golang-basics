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

func (p person) print() {
	fmt.Printf(
		"First Name: %s\nLast Name: %s\nEmail: %s\nZipcode: %d\n\n",
		p.firstName,
		p.lastName,
		p.contact.email,
		p.contact.zipCode,
	)
}

func main() {
	john := person{
		"John",
		"Doe",
		contactInfo{email: "johndoe@doe.com", zipCode: 6784},
	}
	john.print()

	alexContact := contactInfo{email: "alexanderson@anderson.com", zipCode: 6785}
	alex := person{firstName: "Alex", lastName: "Anderson", contact: alexContact}
	alex.print()

	// initialization with "zero-value" as default
	var greg person
	greg.firstName = "Greg"
	greg.lastName = "Peterson"
	greg.contact.email = "gregpeterson@peterson.com"
	greg.contact.zipCode = 6786
	greg.print()
}
