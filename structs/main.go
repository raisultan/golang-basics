package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	john := person{"John", "Doe"}
	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(john.firstName, john.lastName)
	fmt.Println(alex.firstName, alex.lastName)

	// initialization with "zero-value" as default
	var greg person
	fmt.Println("Zero-value initialization", greg)
	greg.firstName = "Greg"
	greg.lastName = "Peterson"
	fmt.Println("Object updated: ", greg.firstName, greg.lastName)
}
