package main

import (
	"testing"
)

func TestNewDeckReturnsFiftySixCards(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
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
