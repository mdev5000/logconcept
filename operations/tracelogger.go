package operations

import (
	"context"
	"github.com/mdev5000/logconcept/attr"
	"github.com/mdev5000/logconcept/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type TraceLogger struct {
	ctx    context.Context
	logger log.Logger
}

func Logger(ctx context.Context) TraceLogger {
	return TraceLogger{
		logger: log.FromContext(ctx),
		ctx:    ctx,
	}
}

func (t TraceLogger) Trace() *Event {
	return t.mkEvent(t.logger.TraceCtx(t.ctx))
}

func (t TraceLogger) Info() *Event {
	return t.mkEvent(t.logger.InfoCtx(t.ctx))
}

func (t TraceLogger) Warn() *Event {
	return t.mkEvent(t.logger.WarnCtx(t.ctx))
}

func (t TraceLogger) Error() *Event {
	return t.mkEvent(t.logger.ErrorCtx(t.ctx))
}

func (t TraceLogger) mkEvent(e *log.Event) *Event {
	return &Event{
		e:     e,
		attrs: attr.KeyValueFromCtx(t.ctx),
		span:  trace.SpanFromContext(t.ctx),
	}
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

func (e *Event) Msg(msg string) {
	e.e.Msg(msg)
	e.span.AddEvent(msg, trace.WithAttributes(e.attrs...))
}
