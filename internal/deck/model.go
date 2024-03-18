package deck

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jvfrodrigues/deck-api/internal/deck/card"
)

var suits = []string{"CLUBS", "DIAMONDS", "HEARTS", "SPADES"}
var values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING", "ACE"}

type Deck struct {
	ID        uuid.UUID   `json:"id"`
	Remaining int         `json:"remaining"`
	Shuffled  bool        `json:"shuffled"`
	Cards     []card.Card `json:"cards"`
}

func NewDeck(shuffled bool) *Deck {
	cards := make([]card.Card, 52)
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, card.Card{
				Value: value,
				Suit:  suit,
				Code:  fmt.Sprintf("%s%s", string(value[0]), string(suit[0])),
			})
		}
	}
	deck := &Deck{
		ID:        uuid.New(),
		Remaining: len(cards),
		Shuffled:  shuffled,
		Cards:     cards,
	}
	return deck
}
