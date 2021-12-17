package log

import "context"

type logContextKey struct {}

func FromContext(ctx context.Context) Logger {
	value := ctx.Value(logContextKey{})
	if value == nil {
		return DefaultLogger()
	}

	logger, ok := value.(Logger)
	if !ok {
		return DefaultLogger()
	}

	return logger
}

func (l Logger) AddToCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, logContextKey{}, l)
}