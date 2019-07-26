package di_test

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

type Zerolog struct {
}

type Config struct {
	Zerolog Zerolog
}

// Printf is a mirror of zerolog's log.Printf.
func (c *Config) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}
