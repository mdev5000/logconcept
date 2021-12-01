package apperror

import (
	"errors"
	"github.com/mdev5000/logconcept/attr"
)

const CodeInternal = 0
const CodeUserError = 400

type AppError struct {
	LogMessage  string
	UserMessage string
	Err         error
	Attr        []attr.Attribute
	Code        int
}

func ToAppError(err error) (ae *AppError) {
	if errors.As(err, &ae) {
		return
	}
	return &AppError{
		UserMessage: "internal error occurred",
		Err:         err,
		Code:        CodeInternal,
	}
}

func (ae *AppError) HttpCode() int {
	if ae.Code == CodeInternal {
		return 500
	}
	return ae.Code
}

func (ae *AppError) IsInternal() bool {
	return ae.Code == CodeInternal
}

func (ae *AppError) Error() string {
	return ae.Err.Error()
}

func InternalErrS(errMsg string, attrs ...attr.Attribute) error {
	return InternalErr(errors.New(errMsg), attrs...)
}

func InternalErr(err error, attrs ...attr.Attribute) error {
	return &AppError{
		Err:  err,
		Attr: attrs,
		Code: CodeInternal,
	}
}

func ExternalErr(code int, userMsg string, err error, attrs ...attr.Attribute) error {
	return &AppError{
		UserMessage: userMsg,
		Err:         err,
		Attr:        attrs,
		Code:        code,
	}
}
