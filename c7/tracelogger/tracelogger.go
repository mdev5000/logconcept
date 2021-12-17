package tracelogger

import (
	"context"
	"github.com/mdev5000/logconcept/c7/log"
	"github.com/mdev5000/logconcept/c7/operations2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type CtxOp = func(ctx context.Context, e Event) Event

type AttributeTraceLogger struct {
	logger *log.Logger
	ctxOp  CtxOp
	span   *trace.Span
}

func NewAttributeTraceLogger(ctx context.Context, ctxOp CtxOp) AttributeTraceLogger {
	span := trace.SpanFromContext(ctx)
	return AttributeTraceLogger{
		logger: operations2.LoggerFromContext(ctx),
		span:   &span,
		ctxOp:  ctxOp,
	}
}

func (l AttributeTraceLogger) Trace() Event {
	return l.mkEvent(l.logger.Trace())
}

func (l AttributeTraceLogger) Info() Event {
	return l.mkEvent(l.logger.Info())
}

func (l AttributeTraceLogger) Warn() Event {
	return l.mkEvent(l.logger.Warn())
}

func (l AttributeTraceLogger) Error() Event {
	return l.mkEvent(l.logger.Warn())
}

func (l AttributeTraceLogger) mkEvent(e *log.Event) Event {
	ev := Event{
		e:    e,
		span: l.span,
	}
	//if l.ctxOp == nil {
	//	return ev
	//}
	//return l.ctxOp(l.ctx, ev)
	return ev
}

type Event struct {
	e     *log.Event
	attrs []attribute.KeyValue
	span  *trace.Span
}

func (e Event) Int(name string, value int) Event {
	return Event{
		e:     e.e.Int(name, value),
		attrs: append(e.attrs, attribute.Int(name, value)),
		span:  e.span,
	}
}

func (e Event) Str(name, value string) Event {
	return Event{
		e:     e.e.Str(name, value),
		attrs: append(e.attrs, attribute.String(name, value)),
		span:  e.span,
	}
}

func (e Event) Err(err error) Event {
	return Event{
		e:     e.e.Err(err),
		attrs: append(e.attrs, attribute.String("error", err.Error())),
		span:  e.span,
	}
}

func (e Event) Msg(msg string) {
	e.e.Msg(msg)
	//e.span.AddEvent(msg, trace.WithAttributes(e.attrs...))
}
