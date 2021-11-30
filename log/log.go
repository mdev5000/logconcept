package log

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Event = zerolog.Event

var defaultLogger Logger

func init() {
	var err error
	defaultLogger, err = NewLogWrapper()
	if err != nil {
		panic(fmt.Errorf("failed to create default log: %w", err))
	}
}

func NewLogWrapper(opts ...Option) (Logger, error) {
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

type CtxOp = func(ctx context.Context, e *Event) *Event

type Logger struct {
	zerolog.Logger
	ctxOp CtxOp
}

func (l Logger) CtxOp() CtxOp {
	return l.ctxOp
}

func (l Logger) WithCtxOp(ctxOp CtxOp) Logger {
	return Logger{
		Logger: l.Logger,
		ctxOp:  ctxOp,
	}
}

func (l Logger) TraceCtx(ctx context.Context) *Event {
	return l.withCtx(ctx, l.Trace())
}

func (l Logger) InfoCtx(ctx context.Context) *Event {
	return l.withCtx(ctx, l.Info())
}

func (l Logger) WarnCtx(ctx context.Context) *Event {
	return l.withCtx(ctx, l.Warn())
}

func (l Logger) ErrorCtx(ctx context.Context) *Event {
	return l.withCtx(ctx, l.Error())
}

func (l Logger) withCtx(ctx context.Context, e *Event) *Event {
	return l.ctxOp(ctx, e)
}
