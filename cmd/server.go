package cmd

import (
	"fmt"
	"net/http"

	"github.com/jvfrodrigues/deck-api/internal/deck"
	"github.com/jvfrodrigues/deck-api/internal/env"
	"github.com/jvfrodrigues/deck-api/internal/logging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func start() error {
	l, _ := logging.NewLogger()

	e := echo.New()
	serverConfig(e, l)
	env := env.LoadEnvVariables(l)

	if env.Debug {
		l.SetLevel(logrus.DebugLevel)
	}

	apiRoute := e.Group("/api")

	deckRoute := apiRoute.Group("/deck")
	handler := deck.NewHandler(env, l)
	handler.UseSubroute(deckRoute)

	apiPort := fmt.Sprintf(":%s", env.APIPort)
	return e.Start(apiPort)
}

func serverConfig(e *echo.Echo, l *logrus.Logger) {
	e.Use(middleware.CORS())
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.SetLevel(log.INFO)
	e.HideBanner = false
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		l.WithField("endpoint", c.Path()).
			WithField("params", c.QueryParams()).
			WithError(err).
			Warn("http error handler")

		he, ok := err.(*echo.HTTPError)
		if !ok {
			he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		e.DefaultHTTPErrorHandler(he, c)
	}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogMethod: true,
			LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
				l.WithFields(logrus.Fields{
					"URI":    values.URI,
					"status": values.Status,
					"method": values.Method,
				}).Info("request")
				return nil
			},
		},
	))
	e.Use(middleware.Recover())
}
