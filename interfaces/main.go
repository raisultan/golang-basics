package main

type englishBot struct {
	text string
}

type spanishBot struct {
	text string
}

func (eb englishBot) getGreeting() string {
	return "Hello, friend!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
