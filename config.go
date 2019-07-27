package di_test

import (
	"github.com/rs/zerolog"
	"io"
	"os"
)

type Zerolog struct {
	// TimeFieldFormat sets the zerolog.TimeFieldFormat
	TimeFieldFormat string
	// GlobalLevel is passed to zerolog.SetGlobalLevel
	GlobalLevel zerolog.Level
	// OutputWriter determines where logs will be written.  Default is os.Stderr.
	OutputWriter io.Writer
}

type Config struct {
	Zerolog Zerolog
	Log     zerolog.Logger
}

// Run sets up objects with configuration given in Config and run necessary background processes.
func (c Config) Run() Config {
	// configure zerolog and set defaults
	{
		// set time format
		if c.Zerolog.TimeFieldFormat == "" {
			zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		}
		zerolog.SetGlobalLevel(c.Zerolog.GlobalLevel)

		// set where output will be written
		if c.Zerolog.OutputWriter == nil {
			c.Zerolog.OutputWriter = os.Stderr
		}
		// create logger with timestamp and caller (line numbers) by default
		c.Log = zerolog.New(c.Zerolog.OutputWriter).With().Timestamp().Logger().With().Caller().Logger()
	}
	return c
}
