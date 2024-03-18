package deck

import "errors"

type DeckRepository struct {
	decks []Deck
}

func NewRepository() *DeckRepository {
	return &DeckRepository{
		decks: make([]Deck, 0),
	}
}

func (r *DeckRepository) Save(deck Deck) error {
	r.decks = append(r.decks, deck)
	return nil
}

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

func (r *DeckRepository) Draw(id string, count int) (*Deck, error) {
	var foundDeck *Deck
	for _, deck := range r.decks {
		if deck.ID.String() == id {
			foundDeck = &deck
		}
	}
	if foundDeck == nil {
		return nil, errors.New("deck not found")
	}
	if count > len(foundDeck.Cards) {
		return nil, errors.New("deck not found")
	}
	for i := 0; i < count; i++ {
		foundDeck.DrawnCards = append(foundDeck.DrawnCards, foundDeck.Cards[0])
		foundDeck.Cards = foundDeck.Cards[:1]
	}
	return foundDeck, nil
}
