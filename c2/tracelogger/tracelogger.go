package tracelogger

import (
	"context"
	"github.com/mdev5000/logconcept/c2/log"
	"github.com/mdev5000/logconcept/c2/operations2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type AttributeTraceLogger struct {
	ctx    context.Context
	logger operations2.Log
	ctxOp  operations2.CtxOp
}

func NewAttributeTraceLogger(ctx context.Context, ctxOp operations2.CtxOp) AttributeTraceLogger {
	return AttributeTraceLogger{
		logger: operations2.LoggerFromContext(ctx),
		ctx:    ctx,
		ctxOp:  ctxOp,
	}
}

func (l AttributeTraceLogger) Trace() log.LogEvent {
	return l.mkEvent(l.logger.Trace())
}

func (l AttributeTraceLogger) Info() log.LogEvent {
	return l.mkEvent(l.logger.Info())
}

func (l AttributeTraceLogger) Warn() log.LogEvent {
	return l.mkEvent(l.logger.Warn())
}

func (l AttributeTraceLogger) Error() log.LogEvent {
	return l.mkEvent(l.logger.Warn())
}

func (l AttributeTraceLogger) mkEvent(e log.LogEvent) log.LogEvent {
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
	e     log.LogEvent
	attrs []attribute.KeyValue
	span  trace.Span
}

func (e *Event) Int(name string, value int) log.LogEvent {
	e.e = e.e.Int(name, value)
	e.attrs = append(e.attrs, attribute.Int(name, value))
	return e
}

func (e *Event) Str(name, value string) log.LogEvent {
	e.e = e.e.Str(name, value)
	e.attrs = append(e.attrs, attribute.String(name, value))
	return e
}

func (e *Event) Err(err error) log.LogEvent {
	e.e = e.e.Err(err)
	e.attrs = append(e.attrs, attribute.String("error", err.Error()))
	return e
}

func (e *Event) Msg(msg string) {
	e.e.Msg(msg)
	e.span.AddEvent(msg, trace.WithAttributes(e.attrs...))
}
