package operations2

import (
	"context"
	"github.com/mdev5000/logconcept/c5/attr"
	"github.com/mdev5000/logconcept/c5/log"
	"go.opentelemetry.io/otel/trace"
)

type CtxOp = func(ctx context.Context, e *log.Event) *log.Event

var NoOp CtxOp = nil

func CtxOpSpan(child CtxOp) CtxOp {
	return func(ctx context.Context, e *log.Event) *log.Event {
		e = LogSpanID(ctx, e)
		if child == nil {
			return e
		}
		return child(ctx, e)
	}
}

func LogSpanID(ctx context.Context, e *log.Event) *log.Event {
	span := trace.SpanFromContext(ctx)
	spanID := ""
	if span != nil {
		spanID = span.SpanContext().TraceID().String()
	}
	e = e.Str("spanID", spanID)
	return e
}

func CtxOpAddAttributes(child CtxOp) CtxOp {
	return func(ctx context.Context, e *log.Event) *log.Event {
		e = AddToLogEvent(e, attr.FromCtx(ctx))
		if child != nil {
			return child(ctx, e)
		}
		return e
	}
}

func AddToLogEvent(e *log.Event, attrs []attr.Attribute) *log.Event {
	for _, a := range attrs {
		e = a.ToEvent(e)
	}
	return e
}
