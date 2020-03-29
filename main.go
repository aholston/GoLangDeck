package main

import "fmt"

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 7)

	// hand.print()
	// remainingDeck.print()

	cards := newDeck()
	fmt.Println(cards.toString())
}
