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
	for _, suit := range cardSuits {
		for _, value := range cardValues {
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

// another receiver for `deck` type
func (d deck) deal(numOfCards int) deck {
	dealtCards := d[:numOfCards]
	return dealtCards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
