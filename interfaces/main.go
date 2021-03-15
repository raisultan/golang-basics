package main

import "fmt"

type bot interface {
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
