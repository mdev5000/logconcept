package tracelogger

import (
	"context"
	"github.com/mdev5000/logconcept/c5/log"
	"github.com/mdev5000/logconcept/c5/operations2"
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
		e:    e,
		span: trace.SpanFromContext(l.ctx),
	}
	if l.ctxOp == nil {
		return ev
	}
	return l.ctxOp(l.ctx, ev)
}

type attr struct {
	kv   attribute.KeyValue
	len  int
	next *attr
}

type Event struct {
	e     *log.Event
	attrs *attr
	span  trace.Span
}

func (e Event) Int(name string, value int) Event {
	e.e = e.e.Int(name, value)
	e.attrs = appendAttr(e.attrs, attribute.Int(name, value))
	return e
}

func appendAttr(existing *attr, kv attribute.KeyValue) *attr {
	length := 1
	if existing != nil {
		length = existing.len
	}
	return &attr{
		kv:   kv,
		len:  length + 1,
		next: existing,
	}
}

func (e Event) Str(name, value string) Event {
	e.e = e.e.Str(name, value)
	e.attrs = appendAttr(e.attrs, attribute.String(name, value))
	return e
}

func (e Event) Err(err error) Event {
	e.e = e.e.Err(err)
	e.attrs = appendAttr(e.attrs, attribute.String("error", err.Error()))
	return e
}

func (e Event) Msg(msg string) {
	e.e.Msg(msg)
	e.span.AddEvent(msg, trace.WithAttributes(collectAttrs(e.attrs)...))
}

func collectAttrs(a *attr) []attribute.KeyValue {
	if a == nil {
		return nil
	}
	kvs := make([]attribute.KeyValue, a.len)
	for i := 0; a != nil; i++ {
		kvs[i] = a.kv
		a = a.next
	}
	return kvs
}
