package userop

import (
	"context"
	"github.com/mdev5000/logconcept/c2/log"
	"github.com/mdev5000/logconcept/c2/operations2"
	"github.com/mdev5000/logconcept/c2/tracelogger"
)

//var ctxOp = operations2.CtxOpSpan(operations2.CtxOpAddAttributes(operations2.NoOp))

type Log = operations2.Log

func ctxOp2(ctx context.Context, e log.LogEvent) log.LogEvent {
	//e = operations2.LogSpanID(ctx, e)
	e.Str("first", "first value")
	e.Str("second", "second value")
	return e
}

func Logger(ctx context.Context) Log {
	return tracelogger.NewAttributeTraceLogger(ctx, nil)
}

func Logger2(ctx context.Context) AppLog {
	return AppLog{tracelogger.NewAttributeTraceLogger(ctx, nil)}
}
