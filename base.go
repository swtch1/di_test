package di_test

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Zerolog struct {
	// TimeFieldFormat sets the zerolog.TimeFieldFormat
	TimeFieldFormat string
	// GlobalLevel is passed to zerolog.SetGlobalLevel
	GlobalLevel int
}

type Config struct {
	Zerolog Zerolog
	Log     Log
}

// Configure sets up objects with configuration given in Config.
func (c *Config) Configure() {
	// set zerolog defaults
	{
		if c.Zerolog.TimeFieldFormat == "" {
			zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		}
		if c.Zerolog.GlobalLevel == 0 {
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		}
	}
}

type Log struct{}

// Printf is a mirror of zerolog's log.Printf.
func (c *Log) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}
