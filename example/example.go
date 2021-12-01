package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/mdev5000/logconcept/apperror"
	"github.com/mdev5000/logconcept/attr"
	"github.com/mdev5000/logconcept/log"
	"github.com/mdev5000/logconcept/operations"
)

func main() {
	fmt.Println("\nLogging example:")

	logger, _ := log.NewLogWrapper()
	logger = logger.WithCtxOp(attr.CtxOpSpan(attr.CtxOpAddAttributes(logger.CtxOp())))

	ctx := context.Background()
	ctx = attr.AddToCtx(ctx, attr.Str("string1", "string1 value"))

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
	err = apperror.InternalErrS("some error occurred",
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
}
