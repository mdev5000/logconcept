package log

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type BaseEvent = zerolog.Event

type Event struct {
	e *BaseEvent
}

func (e *Event) Int(k string, v int) LogEvent {
	e.e = e.e.Int(k, v)
	return e
}

func (e *Event) Err(err error) LogEvent {
	e.e = e.e.Err(err)
	return e
}

func (e *Event) Str(k, v string) LogEvent {
	e.e = e.e.Str(k, v)
	return e
}

func (e *Event) Msg(s string) {
	e.e.Msg(s)
}

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
			return Logger{}, err
		}
	}
	return Logger{Logger: zerolog.New(optsOut.w)}, nil
}

func DefaultLogger() Logger {
	return defaultLogger
}

type CtxOp = func(ctx context.Context, e *BaseEvent) *BaseEvent

type Logger struct {
	zerolog.Logger
}

func (l Logger) Trace() LogEvent {
	return l.wrapEvent(l.Logger.Trace())
}

func (l Logger) Info() LogEvent {
	return l.wrapEvent(l.Logger.Info())
}

func (l Logger) Warn() LogEvent {
	return l.wrapEvent(l.Logger.Warn())
}

func (l Logger) Error() LogEvent {
	return l.wrapEvent(l.Logger.Error())
}

func (l Logger) wrapEvent(e *zerolog.Event) *Event {
	return &Event{e}
}
