package main

import "fmt"

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 7)

	// hand.print()
	// remainingDeck.print()

	cards := getFromFile("My Cards")
	fmt.Println(cards)
}
