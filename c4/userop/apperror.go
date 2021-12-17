package userop

import (
	"github.com/mdev5000/logconcept/c4/apperror"
	"github.com/mdev5000/logconcept/c4/tracelogger"
	"github.com/mdev5000/logconcept/internalerr"
)

type AppLog struct {
	tracelogger.AttributeTraceLogger
}

func (l AppLog) AppError(err error) {
	ae := apperror.ToAppError(err)
	msg := ae.LogMessage

	if msg == "" {
		msg = ae.Err.Error()
	}

	var e tracelogger.Event

	if ae.IsInternal() {
		e = l.Error()
	} else {
		e = l.Info()
	}

	//e = operations2.AddToLogEvent(e, ae.Attr).
	e = e.Err(ae.Err).
		Int("code", ae.Code)

	stack := internalerr.Stack(ae.Err)
	if stack != "" {
		e = e.Str("stack", stack)
	}

	e.Msg(msg)
}
