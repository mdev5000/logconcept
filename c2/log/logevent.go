package log

type LogEvent interface {
	Str(k string, v string) LogEvent
	Int(k string, v int) LogEvent
	Err(err error) LogEvent
	Msg(string)
}
