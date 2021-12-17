package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/mdev5000/logconcept/c5/apperror"
	"github.com/mdev5000/logconcept/c5/attr"
	"github.com/mdev5000/logconcept/c5/log"
	"github.com/mdev5000/logconcept/c5/operations2"
	"github.com/mdev5000/logconcept/c5/userop"
	"github.com/mdev5000/logconcept/internalerr"
	"os"
)

func main() {
	// Setup

	b := bytes.NewBuffer(nil)
	logger, _ := log.New(log.WithWriter(b))

	ctx := context.Background()
	ctx = attr.AddToCtx(ctx,
		attr.Str("string1", "string1 value"))

	ctx = operations2.AddLoggerToCtx(ctx, &logger)

	fmt.Println("\nOperation example:")

	userop.Logger(ctx).Info().
		Str("string2", "string2 value").
		Msg("another message")

	var err error
	fmt.Println("\nFrom errors:")
	err = apperror.InternalErrS(true,
		"some error occurred",
		attr.Str("errStr", "err string value"),
	)
	userop.Logger2(ctx).AppError(err)

	err = apperror.ExternalErr(
		apperror.CodeUserError,
		"cannot do stuff",
		errors.New("this error happened"),
		attr.Int("num fails", 6),
	)
	userop.Logger2(ctx).AppError(err)

	err = internalerr.StackErrF("some error %s", "arg")
	err = apperror.InternalErr(err, attr.Int("someVal", 5))
	userop.Logger2(ctx).AppError(err)

	logData := b.Bytes()
	fmt.Println(string(logData))

	if err := os.WriteFile("example.log", logData, 0775); err != nil {
		panic(err)
	}
}
