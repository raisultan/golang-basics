package main

import (
	"os"
	"testing"
)

const expectedDeckLength = 52

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	deckLength := len(deck)

	if deckLength != expectedDeckLength {
		t.Errorf("Expected deck length of %v, got %v", expectedDeckLength, deckLength)
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", deck[0])
	}

	if deck[deckLength-1] != "Queen of Clubs" {
		t.Errorf("Expected Queen of Clubs, got %v", deck[deckLength-1])
	}
}

func TestWriteToFileAndNewDeckFormFile(t *testing.T) {
	const path = "_decktesting"
	os.Remove(path)

	d := newDeck()
	d.writeToFile(path)
	readDeck := newDeckFromFile(path)
	deckLength := len(readDeck)

	if deckLength != expectedDeckLength {
		t.Errorf("Expected deck length of %v, got %v", expectedDeckLength, deckLength)
	}

	if readDeck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", readDeck[0])
	}

	if readDeck[deckLength-1] != "Queen of Clubs" {
		t.Errorf("Expected Queen of Clubs, got %v", readDeck[deckLength-1])
	}

	os.Remove(path)
}
