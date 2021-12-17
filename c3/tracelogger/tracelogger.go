package tracelogger

import (
	"context"
	"github.com/mdev5000/logconcept/c3/log"
	"github.com/mdev5000/logconcept/c3/operations2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type CtxOp = func(ctx context.Context, e *Event) *Event

type AttributeTraceLogger struct {
	ctx    context.Context
	logger *log.Logger
	ctxOp  CtxOp
}

func NewAttributeTraceLogger(ctx context.Context, ctxOp CtxOp) AttributeTraceLogger {
	return AttributeTraceLogger{
		logger: operations2.LoggerFromContext(ctx),
		ctx:    ctx,
		ctxOp:  ctxOp,
	}
}

func (l AttributeTraceLogger) Trace() *Event {
	return l.mkEvent(l.logger.Trace())
}

func (l AttributeTraceLogger) Info() *Event {
	return l.mkEvent(l.logger.Info())
}

func (l AttributeTraceLogger) Warn() *Event {
	return l.mkEvent(l.logger.Warn())
}

func (l AttributeTraceLogger) Error() *Event {
	return l.mkEvent(l.logger.Warn())
}

func (l AttributeTraceLogger) mkEvent(e *log.Event) *Event {
	ev := &Event{
		e:    e,
		span: trace.SpanFromContext(l.ctx),
	}
	if l.ctxOp == nil {
		return ev
	}
	return l.ctxOp(l.ctx, ev)
}

type Event struct {
	e     *log.Event
	attrs []attribute.KeyValue
	span  trace.Span
}

func (e *Event) Int(name string, value int) *Event {
	e.e = e.e.Int(name, value)
	e.attrs = append(e.attrs, attribute.Int(name, value))
	return e
}

func (e *Event) Str(name, value string) *Event {
	e.e = e.e.Str(name, value)
	e.attrs = append(e.attrs, attribute.String(name, value))
	return e
}

func (e *Event) Err(err error) *Event {
	e.e = e.e.Err(err)
	e.attrs = append(e.attrs, attribute.String("error", err.Error()))
	return e
}

func (e *Event) Msg(msg string) {
	e.e.Msg(msg)
	e.span.AddEvent(msg, trace.WithAttributes(e.attrs...))
}
