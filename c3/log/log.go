package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Event = zerolog.Event

var defaultLogger Logger

func init() {
	var err error
	defaultLogger, err = New()
	if err != nil {
		panic(fmt.Errorf("failed to create default log: %w", err))
	}
}

func New(opts ...Option) (Logger, error) {
	optsOut := options{w: os.Stdout}
	for _, opt := range opts {
		if err := opt(&optsOut); err != nil {
			return zerolog.Logger{}, err
		}
	}
	return zerolog.New(optsOut.w), nil
}

func DefaultLogger() Logger {
	return defaultLogger
}

type Logger = zerolog.Logger
