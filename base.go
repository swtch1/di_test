package di_test

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Zerolog struct {
	TimeFieldFormat string
}


type Config struct {
	Zerolog Zerolog
	Log Log
}

// Configure sets up objects with configuration given in Config.
func (c *Config) Configure() {
	if c.Zerolog.TimeFieldFormat == "" {
		fmt.Println("setting time field format default")
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}
}

type Log struct {}

// Printf is a mirror of zerolog's log.Printf.
func (c *Log) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}
