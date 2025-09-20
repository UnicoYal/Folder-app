package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type LoggerContainer struct {
	logger *zerolog.Logger
}

func New() *LoggerContainer {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &LoggerContainer{
		logger: &logger,
	}
}

func (c *LoggerContainer) ProvideLogger() *zerolog.Logger {
	return c.logger
}
