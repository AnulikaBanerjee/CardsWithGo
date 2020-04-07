package main

import (
	//"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(d))
	}
	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of Clubs as first card but got %v", d[0])
	}
}

func TestSaveToFileReadDeckFromFile(t *testing.T) {
	d := newDeck()
	os.Remove("_deckTestFile")
	d.saveToFile("_deckTestFile")
	savedDecks := readDeckFromFile("_deckTestFile")
	if len(savedDecks) != 52 {
		t.Errorf("Expected 52 cards but got %d", len(savedDecks))
	}
	os.Remove("_deckTestFile")
}
