package main

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 7)

	hand.print()
	remainingDeck.print()
}