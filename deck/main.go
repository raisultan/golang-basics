package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("Full deck: ")
	cards.print()

	dealtCards, leftCards := deal(cards, 5)

	fmt.Println("\n\nDealt cards: ")
	dealtCards.print()

	fmt.Println("\n\nLeft cards: ")
	leftCards.print()

	leftCards.writeToFile("deck")
	readDeck := newDeckFromFile("deck")

	fmt.Println("\n\nRead cards: ")
	readDeck.print()

	readDeck.shuffle()
	fmt.Println("\n\nShuffled cards: ")
	readDeck.print()
}
