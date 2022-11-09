package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// underscores are for variables that we do not care about (rids errors of unused variables)
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of " + cardSuit)
		}
	}
	
	return cards
}

func (cards deck) print () {
	for _, card := range cards {
		fmt.Println(card)
	}
}