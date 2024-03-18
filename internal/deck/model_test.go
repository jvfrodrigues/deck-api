package deck_test

import (
	"testing"

	"github.com/jvfrodrigues/deck-api/internal/deck"
)

func TestNewDeck(t *testing.T) {
	var tests = []struct {
		deck            deck.Deck
		expectedLen     int
		expectedShuffle bool
	}{
		{deck: *deck.NewDeck(), expectedLen: 52, expectedShuffle: false},
		{deck: *deck.NewDeck(deck.Shuffled(true)), expectedLen: 52, expectedShuffle: true},
		{deck: *deck.NewDeck(deck.Partial("AS,AC,AD,AH")), expectedLen: 4, expectedShuffle: false},
		{deck: *deck.NewDeck(deck.Partial("AS,AC,AD,AH"), deck.Shuffled(true)), expectedLen: 4, expectedShuffle: true},
		{deck: *deck.NewDeck(deck.Partial("")), expectedLen: 52, expectedShuffle: false},
	}
	for _, testcase := range tests {
		if len(testcase.deck.Cards) != testcase.expectedLen {
			t.Errorf("got %d, want %d", len(testcase.deck.Cards), testcase.expectedLen)
		}
		if testcase.deck.Shuffled != testcase.expectedShuffle {
			t.Errorf("got %t, want %t", testcase.deck.Shuffled, testcase.expectedShuffle)
		}
	}
}

func TestUnshuffled(t *testing.T) {
	var tests = []struct {
		deck            deck.Deck
		expectFirstCard string
	}{
		{deck: *deck.NewDeck(deck.Shuffled(false)), expectFirstCard: "AS"},
		{deck: *deck.NewDeck(deck.Partial("2S,2D,2C,2H"), deck.Shuffled(false)), expectFirstCard: "2S"},
	}
	for _, testcase := range tests {
		if testcase.deck.Cards[0].Code != testcase.expectFirstCard {
			t.Errorf("got %s, want %s", testcase.deck.Cards[0].Code, testcase.expectFirstCard)
		}
	}
}
