// Package deck provides all requirements for creating and managing decks
package deck

import (
	"net/http"

	"github.com/jvfrodrigues/deck-api/internal/env"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Handler for deck endpoints
type Handler struct {
	logger *logrus.Logger
	env    *env.Vars
}

// UseSubroute registers the deck routes
func (h *Handler) UseSubroute(group *echo.Group) {
	group.POST("/", h.createDeck)
}

func (h *Handler) createDeck(c echo.Context) error {
	deck := NewDeck(false)
	return c.JSON(http.StatusOK, deck)
}

// NewHandler returns a new instance of a deck handler
func NewHandler(
	env *env.Vars,
	logger *logrus.Logger,
) *Handler {
	return &Handler{
		env:    env,
		logger: logger,
	}
}
