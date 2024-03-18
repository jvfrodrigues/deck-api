package deck

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jvfrodrigues/deck-api/internal/deck/card"
)

var suits = []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
var values = []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}

// type Deck represents the structure of a deck
type Deck struct {
	ID         uuid.UUID   `json:"id"`
	Shuffled   bool        `json:"shuffled"`
	Cards      []card.Card `json:"cards"`
	DrawnCards []card.Card `json:"drawn_cards"`
}

// NewDeck creates a new deck
func NewDeck(options ...func(*Deck)) *Deck {
	deck := &Deck{
		ID: uuid.New(),
	}
	fullDeck := fullCardDeck()
	deck.Cards = fullDeck
	deck.DrawnCards = make([]card.Card, 0)
	for _, o := range options {
		o(deck)
	}
	return deck
}

func fullCardDeck() []card.Card {
	cards := make([]card.Card, 0)
	for _, suit := range suits {
		for _, value := range values {
			cardCode := fmt.Sprintf("%s%s", string(value[0]), string(suit[0]))
			cards = append(cards, card.Card{
				Value: value,
				Suit:  suit,
				Code:  cardCode,
			})
		}
	}
	return cards
}

// Shuffled is a option function that alters the creation of the Deck to have its cards shuffled
func Shuffled(shuffle bool) func(*Deck) {
	return func(d *Deck) {
		if shuffle {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := range d.Cards {
				randIndex := rand.Intn(len(d.Cards))
				d.Cards[i], d.Cards[randIndex] = d.Cards[randIndex], d.Cards[i]
			}
			d.Shuffled = shuffle
		}
	}
}

// Partial let's you chose the cards that you want on your deck by passing the card codes
func Partial(cards string) func(*Deck) {
	return func(d *Deck) {
		if cards == "" {
			return
		}
		requestedCards := make(map[string]string)
		for _, card := range strings.Split(strings.ToUpper(cards), ",") {
			requestedCards[card] = card
		}
		partialDeck := make([]card.Card, 0)
		for _, card := range d.Cards {
			if requestedCards[card.Code] != "" {
				partialDeck = append(partialDeck, card)
			}
		}
		d.Cards = partialDeck
	}
}
