package operations

import (
	"github.com/mdev5000/logconcept/apperror"
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

	e.Attrs(ae.Attr).
		Err(ae.Err).
		Int("code", ae.Code).
		Msg(msg)
}