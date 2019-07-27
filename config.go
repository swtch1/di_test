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

type Prometheus struct {
	// Disable will disable the Prometheus endpoint.  Default is false.
	Enable bool
}

type Config struct {
	Zerolog Zerolog
	Prometheus Prometheus
}

type Dependency struct {
	Log     zerolog.Logger
}

// Run sets up objects with configuration given in Config and run necessary background processes.
func (c Config) Run() *Dependency {
	d := &Dependency{}
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
		d.Log = zerolog.New(c.Zerolog.OutputWriter).With().Timestamp().Logger().With().Caller().Logger()
	}

	// configure Prometheus and set defaults, launch endpoint in a goroutine
	{
		if c.Prometheus.Enable {

		}
	}
	return d
}
