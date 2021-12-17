package operations

import (
	"context"
	"github.com/mdev5000/logconcept/c1/attr"
	"github.com/mdev5000/logconcept/c1/log"
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

func (l TraceLogger) Trace() *Event {
	return l.mkEvent(l.logger.TraceCtx(l.ctx))
}

func (l TraceLogger) Info() *Event {
	return l.mkEvent(l.logger.InfoCtx(l.ctx))
}

func (l TraceLogger) Warn() *Event {
	return l.mkEvent(l.logger.WarnCtx(l.ctx))
}

func (l TraceLogger) Error() *Event {
	return l.mkEvent(l.logger.ErrorCtx(l.ctx))
}

func (l TraceLogger) mkEvent(e *log.Event) *Event {
	return &Event{
		e:     e,
		attrs: attr.KeyValueFromCtx(l.ctx),
		span:  trace.SpanFromContext(l.ctx),
	}
}

type Event struct {
	e     *log.Event
	attrs []attribute.KeyValue
	span  trace.Span
}

func (e *Event) Attrs(attrs []attr.Attribute) *Event {
	e.e = attr.AddToLogEvent(e.e, attrs)
	for _, a := range attrs {
		e.attrs = append(e.attrs, a.ToAttribute())
	}
	return e
}

func (e *Event) Attr(a attr.Attribute) *Event {
	e.e = a.ToEvent(e.e)
	e.attrs = append(e.attrs, a.ToAttribute())
	return e
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
