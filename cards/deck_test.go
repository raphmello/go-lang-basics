package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}
	firstCard := deck[0]
	if firstCard != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", firstCard)
	}
	lastCard := deck[len(deck)-1]
	if lastCard != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)
	deck := newDeck()
	deck.saveToFile(fileName)
	loadedDeck := newDeckFromFile(fileName)
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}
	os.Remove(fileName)
}
