package main

func main() {
	cards := newDeck()
	// cards.print()

	hand, remainingDeck := deal(cards, 10)
	hand.print()
	remainingDeck.print()
}
