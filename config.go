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
	// OutputFile, if given, will
}

type Config struct {
	Zerolog Zerolog
	Log     zerolog.Logger
}

// Configure sets up objects with configuration given in Config.
func (c *Config) Configure() *Config {
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
		c.Log = zerolog.New(c.Zerolog.OutputWriter).With().Timestamp().Logger()
	}
	return c
}
