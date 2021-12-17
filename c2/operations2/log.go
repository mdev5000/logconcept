package operations2

import (
	"context"
	"github.com/mdev5000/logconcept/c2/log"
)

type Log interface {
	Trace() log.LogEvent
	Info() log.LogEvent
	Warn() log.LogEvent
	Error() log.LogEvent
}

type logContextKey struct{}

func DefaultLogger() Log {
	return log.DefaultLogger()
}

func LoggerFromContext(ctx context.Context) Log {
	value := ctx.Value(logContextKey{})
	if value == nil {
		return DefaultLogger()
	}

	logger, ok := value.(Log)
	if !ok {
		return DefaultLogger()
	}

	return logger
}

func AddLoggerToCtx(ctx context.Context, l Log) context.Context {
	return context.WithValue(ctx, logContextKey{}, l)
}
