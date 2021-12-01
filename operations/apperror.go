package operations

import (
	"github.com/mdev5000/logconcept/apperror"
	"github.com/mdev5000/logconcept/internalerr"
)

func (l TraceLogger) AppError(err error) {
	ae := apperror.ToAppError(err)
	msg := ae.LogMessage

	if msg == "" {
		msg = ae.Err.Error()
	}

	var e *Event

	if ae.IsInternal() {
		e = l.Error()
	} else {
		e = l.Info()
	}

	e = e.Attrs(ae.Attr).
		Err(ae.Err).
		Int("code", ae.Code)

	stack := internalerr.Stack(ae.Err)
	if stack != "" {
		e = e.Str("stack", stack)
	}

	e.Msg(msg)
}