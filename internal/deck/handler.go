// Package deck provides all requirements for creating and managing decks
package deck

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/jvfrodrigues/deck-api/internal/deck/card"
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
	group.POST("", h.createDeck)
	group.GET("/:id", h.openDeck)
	group.GET("/:id/draw/:count", h.drawDeck)
}

type createDeckResponse struct {
	ID        uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
}

// CreateDeck godoc
//
//	@Summary		Creates a new deck
//	@Description	Creates a new deck that can be partial and/or shuffled
//	@Tags			decks
//	@Param			shuffled	query	string	false	"Indicate if deck must be shuffled"
//	@Param			cards		query	string	false	"Give cards wanted on deck"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	createDeckResponse
//	@Failure		500	{object}	map[string]string
//	@Router			/deck/ [post]
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

type getDeckResponse struct {
	ID        uuid.UUID   `json:"deck_id"`
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
	Cards     []card.Card `json:"cards"`
}

// OpenDeck godoc
//
//	@Summary		Shows existing deck
//	@Description	Gets a deck by its ID and shows remaining cards
//	@Tags			decks
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"DeckID"
//	@Success		200	{object}	getDeckResponse
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/deck/{id} [get]
func (h *Handler) openDeck(c echo.Context) error {
	deckID := c.Param("id")
	deck, err := h.repository.Get(deckID)
	if err != nil {
		return err
	}
	response := &getDeckResponse{
		ID:        deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
		Cards:     deck.Cards,
	}
	return c.JSON(http.StatusOK, response)
}

type drawDeckResponse struct {
	Cards []card.Card `json:"cards"`
}

// DrawDeck godoc
//
//	@Summary		Draws cards from deck
//	@Description	Gets a deck by its ID and draws the amount of cards requested
//	@Tags			decks
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"DeckID"
//	@Param			count	path		string	true	"Card count to draw"
//	@Success		200		{object}	drawDeckResponse
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/deck/{id}/draw/{count} [get]
func (h *Handler) drawDeck(c echo.Context) error {
	deckID := c.Param("id")
	strCount := c.Param("count")
	if deckID == "" || strCount == "" {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "missing required params, deckID and count"})
	}
	count, err := strconv.Atoi(strCount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "count must be a number"})
	}
	drawnCards, err := h.repository.Draw(deckID, count)
	if err != nil {
		return err
	}
	response := &drawDeckResponse{
		Cards: drawnCards,
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
