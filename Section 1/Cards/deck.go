package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// underscores are for variables that we do not care about (rids errors of unused variables)
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of " + cardSuit)
		}
	}
	
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	deckString := strings.Join([]string(d), ",")

	return deckString
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


func (d deck) saveToFile(fileName string) error {
	deckString := d.toString()
	error := os.WriteFile(fileName, []byte(deckString), 0666)
	return error

}

func newDeckFromFile(fileName string) (deck) {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)

}