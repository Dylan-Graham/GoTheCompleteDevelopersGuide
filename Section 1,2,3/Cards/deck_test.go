package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Error("Expected deck of length 16, but got ", len(deck))
	}

	if deck[0] != "Ace of Hearts" {
		t.Errorf("Expected first element in the new deck to be 'Ace of Hearts', but got '%s'", deck[0])
	}

	lastIndex := len(deck) - 1
	if deck[lastIndex] != "Four of Clubs" {
		t.Errorf("Expected last element in the new deck to be 'Four of Clubs', but got '%s'", deck[lastIndex])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "_testDeck"
	os.Remove(testFileName)

	deck := newDeck()

	if len(deck) != 16 {
		t.Error("Expected deck of length 16, but got ", len(deck))
	}

	error := deck.saveToFile(testFileName)

	if error != nil {
		t.Error("Did not expect an error during the save of a file")
	}

	deckFromFile := newDeckFromFile(testFileName)

	if len(deckFromFile) != 16 {
		t.Error("Expected deck of length 16, but got ", len(deckFromFile))
	}

	os.Remove(testFileName)
}