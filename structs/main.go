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

// "&" and "*" operators
// "&" before variable gives address of the variable in the memory
// "*" before the address of a variable in the memory gives its value that lies on that address
// "*type" is not pointer, instead it is a type description that says: "pointer to this type"

func (p *person) updateFirstName(newFirstName string) {
	// get actual value from address and adjust it
	(*p).firstName = newFirstName
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

	// call updateFirstName() on value pointer
	// ideally it must written like this (&greg).updateFirstName("Gregory")
	// but compiler can differentiate and correctly get pointer or value if type of parameter in the function is pointer type
	greg.updateFirstName("Gregory")
	greg.print()
}

// gotcha moment: with reference types u don't need to pass address, it is itself is an address
// and the change that made to value that is reference type directly changes the value, with no need of passing address

// reference types is a complex data type that looks like wrapper data + value
// and the wrapper data has pointer to value and when you assign or get reference type, what you are actually receiving is wrapper that already keeps pointer

// VALUE TYPES: int, float, string, bool, struct - use pointers
// REFERENCE TYPES: slice, map, channel, pointer, function - do not use pointers
