package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected card Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected cardFour of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	newDeck := newDeckFromFile("_decktesting")

	if newDeck == nil {
		t.Errorf("Expected to load deck from file, but got nil")
	}

	if len(newDeck) != 16 {
		t.Errorf("Expected new deck length of 16, but got %v", len(newDeck))
	}

	if newDeck[0] != "Ace of Spades" {
		t.Errorf("Expected card Ace of Spades from new deck, but got %v", newDeck[0])
	}

	if newDeck[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected cardFour of Clubs from new deck, but got %v", newDeck[len(newDeck)-1])
	}

	os.Remove("_decktesting")
}
