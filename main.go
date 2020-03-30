package main

func main() {
	cards := getFromFile("My Cards")
	cards.shuffle()
	cards.print()
}
