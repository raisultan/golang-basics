package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // property name may be left empty then, struct name will be the property name - contactInfo in our case
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

// golang is "pass by value" language. it means that when executing a function,
// compiler makes a copy of value and then passes it to the function body
// thus, on things like adjusting a value we must pass not a value but a pointer to that value

func (personPointer *person) updateFirstName(newFirstName string) {
	(*personPointer).firstName = newFirstName
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

	gregPointer := &greg
	gregPointer.updateFirstName("Gregory")
	greg.print()
}
