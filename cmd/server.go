package cmd

import (
	"fmt"
	"net/http"

	"github.com/jvfrodrigues/deck-api/internal/env"
	"github.com/jvfrodrigues/deck-api/internal/logging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

func start() error {
	l, _ := logging.NewLogger()

	e := echo.New()
	serverConfig(e, l)
	env := env.LoadEnvVariables(l)

	if env.Debug {
		l.SetLevel(logrus.DebugLevel)
	}

	apiPort := fmt.Sprintf("Server running on :%s", env.APIPort)
	return e.Start(apiPort)
}

func serverConfig(e *echo.Echo, l *logrus.Logger) {
	e.Use(middleware.CORS())
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
	e.Use(contentTypeSetMiddleWare())
}

func contentTypeSetMiddleWare() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().
				Header().
				Set(echo.HeaderContentType, "application/json; schema=object; version=1; cache=none")
			return next(c)
		}
	}
}
