// Package deck provides all requirements for creating and managing decks
package deck

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jvfrodrigues/deck-api/internal/env"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Handler for deck endpoints
type Handler struct {
	logger     *logrus.Logger
	env        *env.Vars
	repository *DeckRepository
}

// UseSubroute registers the deck routes
func (h *Handler) UseSubroute(group *echo.Group) {
	group.POST("/", h.createDeck)
}

type createDeckResponse struct {
	ID        uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
}

func (h *Handler) createDeck(c echo.Context) error {
	shuffled := c.QueryParam("shuffled") == "true"
	cards := c.QueryParam("cards")
	deck := NewDeck(Shuffled(shuffled), Partial(cards))
	err := h.repository.Save(*deck)
	if err != nil {
		return err
	}
	response := &createDeckResponse{
		ID:        deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
	}
	return c.JSON(http.StatusOK, response)
}

// NewHandler returns a new instance of a deck handler
func NewHandler(
	env *env.Vars,
	logger *logrus.Logger,
) *Handler {
	repository := NewRepository()
	return &Handler{
		env:        env,
		logger:     logger,
		repository: repository,
	}
}
