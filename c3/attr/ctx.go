package attr

import (
	"context"
)

type attributeCtxKey struct {}

type attributeStore struct {
	values []Attribute
}

func (as *attributeStore) Push(attr Attribute) {
	as.values = append(as.values, attr)
}

func (as *attributeStore) Values() []Attribute {
	return as.values
}

func AddToCtx(ctx context.Context, attrs ...Attribute) context.Context {
	store, ok := ctx.Value(attributeCtxKey{}).(*attributeStore)
	if !ok || store == nil {
		store = &attributeStore{}
		ctx = context.WithValue(ctx, attributeCtxKey{}, store)
	}
	for _, attr := range attrs {
		store.Push(attr)
	}
	return ctx
}

func FromCtx(ctx context.Context) []Attribute {
	store, ok := ctx.Value(attributeCtxKey{}).(*attributeStore)
	if !ok || store == nil {
		return nil
	}
	return store.Values()
}
