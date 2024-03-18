package deck_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jvfrodrigues/deck-api/internal/deck"
)

func TestCreateAndFindRepository(t *testing.T) {
	existingDeck := deck.NewDeck()
	var tests = []struct {
		givenDeck     deck.Deck
		findID        uuid.UUID
		expectedError error
	}{
		{givenDeck: *existingDeck, findID: existingDeck.ID, expectedError: nil},
		{givenDeck: *existingDeck, findID: uuid.New(), expectedError: errors.New("deck not found")},
	}
	for _, testcase := range tests {
		repository := deck.NewRepository()
		repository.Save(testcase.givenDeck)
		_, err := repository.Get(testcase.findID.String())
		if err != testcase.expectedError {
			if err.Error() != testcase.expectedError.Error() {
				t.Errorf("got %s, want %s", err.Error(), testcase.expectedError.Error())
			}
		}
	}
}

func TestCreateAndDrawRepository(t *testing.T) {
	existingDeck := deck.NewDeck()
	existingPartialDeck := deck.NewDeck(deck.Partial("AS,AH"))
	var tests = []struct {
		givenDeck         deck.Deck
		findID            uuid.UUID
		drawCount         int
		expectedDrawCount int
		expectedFirstCard string
		expectedError     error
	}{
		{givenDeck: *existingDeck, findID: existingDeck.ID, drawCount: 3, expectedDrawCount: 3, expectedFirstCard: "4S", expectedError: nil},
		{givenDeck: *existingDeck, findID: uuid.New(), drawCount: 0, expectedDrawCount: 0, expectedFirstCard: "", expectedError: errors.New("deck not found")},
		{givenDeck: *existingPartialDeck, findID: existingPartialDeck.ID, drawCount: 5, expectedDrawCount: 0, expectedFirstCard: "AS", expectedError: errors.New("not enough cards to draw")},
	}
	for _, testcase := range tests {
		repository := deck.NewRepository()
		repository.Save(testcase.givenDeck)
		cards, err := repository.Draw(testcase.findID.String(), testcase.drawCount)
		if err != testcase.expectedError {
			if err.Error() != testcase.expectedError.Error() {
				t.Errorf("got %s, want %s", err.Error(), testcase.expectedError.Error())
			}
		}
		deck, err := repository.Get(testcase.findID.String())
		if err != nil {
			if err != testcase.expectedError {
				if err.Error() != testcase.expectedError.Error() {
					t.Errorf("got %s, want %s", err.Error(), testcase.expectedError.Error())
				}
			}
		} else {
			if len(cards) != testcase.expectedDrawCount {
				t.Errorf("got %d, want %d", len(cards), testcase.expectedDrawCount)
			}
			if deck.Cards[0].Code != testcase.expectedFirstCard {
				t.Errorf("got %s, want %s", deck.Cards[0].Code, testcase.expectedFirstCard)
			}
		}
	}
}
