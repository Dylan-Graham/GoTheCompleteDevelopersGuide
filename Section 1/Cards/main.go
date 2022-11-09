package main

import (
	"fmt"
)

func main() {
	card := newCard()
	fmt.Println(card)

	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Two of Hearts") 
	// NB: append returns a new slice
	
	cards.print()

	freshDeck := newDeck()
	freshDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}