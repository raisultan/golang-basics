package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

// can omit the parameter name in receiver funcs, if param is not used in func itself
func (englishBot) getGreeting() string {
	return "Hello, friend!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// golang does not support overloading, os u can not have two funcs with the same name
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
