package di_test

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Zerolog struct {
}


type Config struct {
	Zerolog Zerolog
	Log Log
}

// Configure sets up objects with configuration given in Config.
func (c *Config) Configure() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

type Log struct {}

// Printf is a mirror of zerolog's log.Printf.
func (c *Log) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}
