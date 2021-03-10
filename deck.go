package main

import "fmt"

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Solider",
		"King",
		"Queen",
	}

	// empty deck
	cards := deck{}

	// fill the deck with cards
	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// `(d deck)` is a receiver func for type deck - use: deck.print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
