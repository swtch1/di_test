package di_test

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Log struct{}

// Print is a mirror of zerolog's log.Print().
func (_ *Log) Print(v ...interface{}) {
	log.Print(v)
}

// Printf is a mirror of zerolog's log.Printf().
func (_ *Log) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

// Debug is a mirror of zerolog's log.Debug().
func (_ *Log) Debug() *zerolog.Event {
	return log.Debug()
}

// Info is a mirror of zerolog's log.Info().
func (_ *Log) Info() *zerolog.Event {
	return log.Info()
}

// Warn is a mirror of zerolog's log.Warn().
func (_ *Log) Warn() *zerolog.Event {
	return log.Warn()
}

// Error is a mirror of zerolog's log.Error().
func (_ *Log) Error() *zerolog.Event {
	return log.Error()
}

// Fatal is a mirror of zerolog's log.Fatal().
func (_ *Log) Fatal() *zerolog.Event {
	return log.Fatal()
}
