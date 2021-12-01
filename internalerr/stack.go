package internalerr

import (
	"fmt"
	"github.com/pkg/errors"
)

func Stack(err error) string {
	var se WithStack
	if errors.As(err, &se) {
		return se.Stack()
	}
	return ""
}

type WithStack struct {
	err error
}

func (e WithStack) InnerError() error {
	return e.err
}

func (e WithStack) Error() string {
	return e.err.Error()
}

func (e WithStack) Stack() string {
	return fmt.Sprintf("%+v", e.err)
}

func (e WithStack) Unwrap() error { return e.err }


func StackErr(err error) error {
	return WithStack{err: errors.WithStack(err)}
}

func StackErrF(msg string, args ...interface{}) error {
	return StackErr(fmt.Errorf(msg, args...))
}