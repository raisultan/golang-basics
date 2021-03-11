package main

import "fmt"

func printColors(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

func printChoices(c map[string]int) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}
	printColors(colors)

	var movies map[string]int
	fmt.Println(movies)

	choices := make(map[string]int)
	choices["idk"] = -1
	choices["No"] = 0
	choices["Yes"] = 1

	delete(choices, "idk")
	printChoices(choices)
}
