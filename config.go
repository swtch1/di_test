package di_test

import (
	"github.com/rs/zerolog"
)

type Zerolog struct {
	// TimeFieldFormat sets the zerolog.TimeFieldFormat
	TimeFieldFormat string
	// GlobalLevel is passed to zerolog.SetGlobalLevel
	GlobalLevel zerolog.Level
}

type Config struct {
	Zerolog Zerolog
	Log     Log
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
