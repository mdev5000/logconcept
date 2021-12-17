package userop

import (
	"context"
	"github.com/mdev5000/logconcept/c5/tracelogger"
)

//var ctxOp = operations2.CtxOpSpan(operations2.CtxOpAddAttributes(operations2.NoOp))

type Event = tracelogger.Event

type Log = tracelogger.AttributeTraceLogger

func ctxOp2(ctx context.Context, e Event) Event {
	//e = operations2.LogSpanID(ctx, e)
	e.Str("first", "first value")
	e.Str("second", "second value")
	return e
}

func Logger(ctx context.Context) Log {
	return tracelogger.NewAttributeTraceLogger(ctx, nil)
}

func Logger2(ctx context.Context) AppLog {
	return AppLog{tracelogger.NewAttributeTraceLogger(ctx, ctxOp2)}
}
