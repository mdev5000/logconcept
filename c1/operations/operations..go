package operations

import "context"

type Op struct {
}

func (o Op) Finish(err error) {
}

func New(ctx context.Context, name string) (Op, context.Context) {
	return Op{}, ctx
}
