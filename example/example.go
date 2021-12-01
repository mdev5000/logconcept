package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/mdev5000/logconcept/apperror"
	"github.com/mdev5000/logconcept/attr"
	"github.com/mdev5000/logconcept/internalerr"
	"github.com/mdev5000/logconcept/log"
	"github.com/mdev5000/logconcept/operations"
	"os"
)

func main() {
	fmt.Println("\nLogging example:")

	b := bytes.NewBuffer(nil)
	logger, _ := log.New(log.WithWriter(b))
	logger = logger.WithCtxOp(attr.CtxOpSpan(attr.CtxOpAddAttributes(logger.CtxOp())))

	ctx := context.Background()
	ctx = attr.AddToCtx(ctx,
		attr.Str("string1", "string1 value"))

	logger.InfoCtx(ctx).
		Str("another", "value").
		Msg("some message")


	ctx = logger.AddToCtx(ctx)

	fmt.Println("\nOperation example:")
	op, ctx := operations.New(ctx, "my_operation")
	var err error
	defer func() {
		op.Finish(err)
	}()

	operations.Logger(ctx).Info().
		Str("string2", "string2 value").
		Msg("another message")

	fmt.Println("\nFrom errors:")
	err = apperror.InternalErrS(true,
		"some error occurred",
		attr.Str("errStr", "err string value"),
	)
	operations.Logger(ctx).AppError(err)

	err = apperror.ExternalErr(
		apperror.CodeUserError,
		"cannot do stuff",
		errors.New("this error happened"),
		attr.Int("num fails", 6),
	)
	operations.Logger(ctx).AppError(err)

	err = internalerr.StackErrF("some error %s", "arg")
	err = apperror.InternalErr(err, attr.Int("someVal", 5))
	operations.Logger(ctx).AppError(err)

	logData := b.Bytes()
	fmt.Println(string(logData))

	if err := os.WriteFile("example.log", logData, 0775); err != nil {
		panic(err)
	}
}
