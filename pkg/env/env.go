// Package env contains methods related to accessing environment variables
package env

import (
	"flag"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Vars contains the environment variable values
type Vars struct {
	Debug     bool
	APIPort   string
	IsTesting bool
}

// LoadEnvVariables loads the environment variables into the structure and returns them
func LoadEnvVariables(logger *logrus.Logger) *Vars {
	debug := false
	debugStr := os.Getenv("DEBUG")
	debugStr = strings.ToLower(debugStr)
	if strings.ToLower(debugStr) == "true" {
		debug = true
	}

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = "8484"
		logger.WithField("apiPort", apiPort).
			Warn("empty API_PORT env variable. Using default")
	}

	isTesting := false
	if flag.Lookup("test.v") != nil {
		isTesting = true
	}

	logger.WithFields(logrus.Fields{
		"debug":     debug,
		"apiPort":   apiPort,
		"isTesting": isTesting,
	}).Info("environment variables loaded")

	return &Vars{
		Debug:     debug,
		APIPort:   apiPort,
		IsTesting: isTesting,
	}
}
