package attr

import (
	"context"
	"github.com/mdev5000/logconcept/c1/log"
	"go.opentelemetry.io/otel/trace"
)

func CtxOpSpan(child log.CtxOp) log.CtxOp {
	return func(ctx context.Context, e *log.Event) *log.Event {
		span := trace.SpanFromContext(ctx)
		spanID := ""
		if span == nil {
			spanID = span.SpanContext().TraceID().String()
		}
		e = e.Str("spanID", spanID)
		if child != nil {
			return child(ctx, e)
		}
		return e
	}
}

func CtxOpAddAttributes(child log.CtxOp) log.CtxOp {
	return func(ctx context.Context, e *log.Event) *log.Event {
		e = AddToLogEvent(e, FromCtx(ctx))
		if child != nil {
			return child(ctx, e)
		}
		return e
	}
}

func AddToLogEvent(e *log.Event, attrs []Attribute) *log.Event {
	for _, a := range attrs {
		e = a.ToEvent(e)
	}
	return e
}