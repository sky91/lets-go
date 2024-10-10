package dox

import (
	"context"
	"github.com/samber/do/v2"
)

func WithValue[T any](ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, do.NameOf[T](), value)
}

func GetValue[T any](ctx context.Context) (value T, ok bool) {
	value, ok = ctx.Value(do.NameOf[T]()).(T)
	return
}
