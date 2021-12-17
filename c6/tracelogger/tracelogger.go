package tracelogger

import (
	"context"
	"github.com/mdev5000/logconcept/c6/log"
	"github.com/mdev5000/logconcept/c6/operations2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type CtxOp = func(ctx context.Context, e Event) Event

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
		e:     e,
		attri: 0,
		span:  trace.SpanFromContext(l.ctx),
	}
	if l.ctxOp == nil {
		return ev
	}
	return l.ctxOp(l.ctx, ev)
}

type Event struct {
	e     *log.Event
	attri int
	attrs [10]attribute.KeyValue
	span  trace.Span
}

func (e Event) Int(name string, value int) Event {
	e.e = e.e.Int(name, value)
	return e.pushKv(attribute.Int(name, value))
}

func (e Event) pushKv(kv attribute.KeyValue) Event {
	if e.attri == 9 {
		e.attri = 0
	}
	e.attrs[e.attri] = kv
	return e
}

func (e Event) Str(name, value string) Event {
	e.e = e.e.Str(name, value)
	return e.pushKv(attribute.String(name, value))
}

func (e Event) Err(err error) Event {
	e.e = e.e.Err(err)
	return e.pushKv(attribute.String("error", err.Error()))
}

func (e Event) Msg(msg string) {
	e.e.Msg(msg)
	e.span.AddEvent(msg, trace.WithAttributes(e.attrs[:]...))
}
