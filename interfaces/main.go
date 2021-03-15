package main

import "fmt"

// declaring common behavior that bots wil have
// so if any custom type has fields and methods of bot - interface
// they become member of type bot
type bot interface {
	// any custom type that has receiver - getGreeting() is available
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// can omit the parameter name in receiver funcs, if param is not used in func itself
func (englishBot) getGreeting() string {
	return "Hello, friend!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
