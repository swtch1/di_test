package di_test

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

type Log struct {}

func (l *Log) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}
