package main

import (
	"context"
	"github.com/mdev5000/logconcept/attr"
	"github.com/mdev5000/logconcept/log"
	"github.com/mdev5000/logconcept/operations"
)

func main() {
	logger, _ := log.NewLogWrapper()
	logger = logger.WithCtxOp(attr.CtxOpSpan(attr.CtxOpAddAttributes(logger.CtxOp())))

	ctx := context.Background()
	ctx = attr.AddToCtx(ctx, attr.Str("string1", "string1 value"))
	ctx = logger.AddToCtx(ctx)

	logger.InfoCtx(ctx).
		Str("another", "value").
		Msg("some message")


	ctx = operations.New(ctx, "my_operation")

	operations.Logger(ctx).Info().
		Str("string2", "string2 value").
		Msg("another message")
}
