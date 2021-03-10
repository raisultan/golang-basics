package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	deckString := strings.Join([]string(d), "\n")
	return deckString
}

func (d deck) writeToFile(path string) {
	deckString := d.toString()
	content := []byte(deckString)
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeckFromFile(path string) deck {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	contentString := string(content)
	cards := strings.Split(contentString, "\n")
	return deck(cards)
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i int, j int) { d[i], d[j] = d[j], d[i] })
}
