package operations2

import (
	"context"
	"github.com/mdev5000/logconcept/c6/log"
)

type Logger = log.Logger

type logContextKey struct{}

func DefaultLogger() Logger {
	return log.DefaultLogger()
}

func LoggerFromContext(ctx context.Context) *Logger {
	value := ctx.Value(logContextKey{})
	if value == nil {
		l := DefaultLogger()
		return &l
	}

	logger, ok := value.(*Logger)
	if !ok {
		// @todo review this
		l := DefaultLogger()
		return &l
	}

	return logger
}

func AddLoggerToCtx(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, logContextKey{}, l)
}
