package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if actualLen := len(d); actualLen != 52 {
		t.Errorf("Expected deck length of 56, got %v", actualLen)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", d[0])
	}

	if d[51] != "Queen of Clubs" {
		t.Errorf("Expected Queen of Clubs, got %v", d[0])
	}
}

func TestWriteToFileAndNewDeckFormFile(t *testing.T) {
	path := "_decktesting"
	os.Remove(path)

	d := newDeck()
	d.writeToFile(path)
	readDeck := newDeckFromFile(path)

	if actualLen := len(readDeck); actualLen != 52 {
		t.Errorf("Expected deck length of 56, got %v", actualLen)
	}

	if readDeck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", readDeck[0])
	}

	if readDeck[51] != "Queen of Clubs" {
		t.Errorf("Expected Queen of Clubs, got %v", readDeck[0])
	}

	os.Remove(path)
}
