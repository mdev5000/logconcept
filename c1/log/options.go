package log

import (
	"errors"
	"io"
)

type options struct {
	w io.Writer
}
type Option = func(opts *options) error

func WithWriter(w io.Writer) Option {
	return func(opts *options) error {
		if w == nil {
			return errors.New("cannot set Writer to nil")
		}
		opts.w = w
		return nil
	}
}
