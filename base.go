package di_test

import (
	"fmt"
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
func (c *Config) Configure() *Config {
	// set zerolog defaults
	{
		if c.Zerolog.TimeFieldFormat == "" {
			zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		}
		if c.Zerolog.GlobalLevel == 0 {
			fmt.Println("setting global log level")  // FIXME: testing
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		}
	}
	return c
}

type Log struct{}

// Printf is a mirror of zerolog's log.Printf.
func (_ *Log) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

func (_ *Log) Warn(s string) {
	log.Warn().Msg(s)
}
