package deck

import (
	"errors"

	"github.com/jvfrodrigues/deck-api/internal/deck/card"
)

// DeckRepository has the actual implementation of the deck repository
type DeckRepository struct {
	decks []Deck
}

// NewRepository creates a new DeckRepository
func NewRepository() *DeckRepository {
	return &DeckRepository{
		decks: make([]Deck, 0),
	}
}

// Save adds a new deck to the DeckRepository
func (r *DeckRepository) Save(deck Deck) error {
	r.decks = append(r.decks, deck)
	return nil
}

// Get gets a deck from the DeckRepository by its ID
func (r *DeckRepository) Get(id string) (*Deck, error) {
	var foundDeck *Deck
	for _, deck := range r.decks {
		if deck.ID.String() == id {
			foundDeck = &deck
		}
	}
	if foundDeck == nil {
		return nil, errors.New("deck not found")
	}
	return foundDeck, nil
}

// Draw gets an existing deck from the DeckRepository and draws the number of cards passed
func (r *DeckRepository) Draw(id string, count int) ([]card.Card, error) {
	var foundDeck *Deck
	for i, deck := range r.decks {
		if deck.ID.String() == id {
			foundDeck = &r.decks[i]
			break
		}
	}
	if foundDeck == nil {
		return nil, errors.New("deck not found")
	}
	if count > len(foundDeck.Cards) {
		return nil, errors.New("not enough cards to draw")
	}
	drawnCards := make([]card.Card, count)
	for i := 0; i < count; i++ {
		foundDeck.DrawnCards = append(foundDeck.DrawnCards, foundDeck.Cards[0])
		drawnCards[i] = foundDeck.Cards[0]
		foundDeck.Cards = foundDeck.Cards[1:]
	}
	return drawnCards, nil
}
