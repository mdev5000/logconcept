package attr

import (
	"github.com/mdev5000/logconcept/c3/log"
	"go.opentelemetry.io/otel/attribute"
)

type Event = *log.Event

type Attribute interface {
	ToEvent(logEvent Event) Event
	ToAttribute() attribute.KeyValue
}

type AttributeVal struct {
	eventFn func(e Event) Event
	traceFn func() attribute.KeyValue
}

func (a AttributeVal) ToEvent(logEvent Event) Event {
	return a.eventFn(logEvent)
}

func (a AttributeVal) ToAttribute() attribute.KeyValue {
	return a.traceFn()
}

func Str(key, value string) Attribute {
	return AttributeVal{
		eventFn: func(e Event) Event {
			return e.Str(key, value)
		},
		traceFn: func() attribute.KeyValue {
			return attribute.String(key, value)
		},
	}
}

func Int(key string, value int) Attribute {
	return AttributeVal{
		eventFn: func(e Event) Event {
			return e.Int(key, value)
		},
		traceFn: func() attribute.KeyValue {
			return attribute.Int(key, value)
		},
	}
}

//func InterfaceAttr(key string, value interface{}) Attribute {
//	return AttributeVal{
//		eventFn: func(e Event) Event {
//			return e.Interface(key, value)
//		},
//		traceFn: func() attribute.KeyValue {
//			b, _ := json.Marshal(value)
//			return attribute.String(key, string(b))
//		},
//	}
//}
