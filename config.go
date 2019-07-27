package di_test

import (
	"github.com/rs/zerolog"
	"os"
)

type Zerolog struct {
	// TimeFieldFormat sets the zerolog.TimeFieldFormat
	TimeFieldFormat string
	// GlobalLevel is passed to zerolog.SetGlobalLevel
	GlobalLevel zerolog.Level
}

type Config struct {
	Zerolog Zerolog
	Log     zerolog.Logger
}

// Configure sets up objects with configuration given in Config.
func (c *Config) Configure() *Config {
	// set zerolog defaults
	{
		if c.Zerolog.TimeFieldFormat == "" {
			zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		}
		zerolog.SetGlobalLevel(c.Zerolog.GlobalLevel)
	}
	return c
}

// Logger is the global logger.
var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
