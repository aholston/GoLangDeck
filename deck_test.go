package main

import (
	"os"
	"testing"
)

func TestNewDeckReturnsFiftySixCards(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	}
}

func TestNewDeckFirstAndLastCard(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card to be King of Clubs, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := getFromFile("_decktesting")

	if len(loadedDeck) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
