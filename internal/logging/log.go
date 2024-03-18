package logging

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

// NewLogger creates a new instance of the logrus logger with our defined formatting
func NewLogger() (*log.Logger, error) {
	logger := log.New()

	logger.SetOutput(io.Discard) // Send all logs to nowhere by default

	logger.SetFormatter(&log.JSONFormatter{})

	logger.SetLevel(log.InfoLevel)

	logger.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})
	logger.AddHook(&writer.Hook{ // Send info and debug logs to stdout
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.InfoLevel,
			log.DebugLevel,
		},
	})
	return logger, nil
}
