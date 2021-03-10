package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("Ful deck: ")
	cards.print()

	dealtCards, leftCards := deal(cards, 5)

	fmt.Println("\n\nDealt cards: ")
	dealtCards.print()

	fmt.Println("\n\nLeft cards: ")
	leftCards.print()
}
