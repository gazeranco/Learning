// test file

package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("expected deck length of 52 but got: %v", len(d))
	}

	//ace of spaces 1st
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got: %v", d[0])
	}

	//king of clubs last
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs but got: %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("expected deck length of 52 but got: %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
