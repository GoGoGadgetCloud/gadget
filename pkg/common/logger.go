package common

import (
	"os"

	"github.com/charmbracelet/log"
)

type (
	DefaultLogger struct {
		err *log.Logger
		out *log.Logger
	}

	Logger interface {
		Err() *log.Logger
		Out() *log.Logger
	}
)

func NewDefaultLogger() Logger {
	return &DefaultLogger{
		err: log.NewWithOptions(os.Stderr, log.Options{
			ReportCaller:    false,
			ReportTimestamp: false,
			TimeFormat:      log.DefaultTimeFormat,
		}),
		out: log.NewWithOptions(os.Stdout, log.Options{
			ReportCaller:    true,
			ReportTimestamp: false,
			TimeFormat:      log.DefaultTimeFormat,
		}),
	}
}

func (d *DefaultLogger) Err() *log.Logger {
	return d.err
}

func (d *DefaultLogger) Out() *log.Logger {
	return d.out
}
