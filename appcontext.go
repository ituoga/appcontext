package appcontext

import "context"

type Key[T any] string

func With[T any](ctx context.Context, k Key[T], v T) context.Context {
	return context.WithValue(ctx, k, v)
}

func Get[T any](ctx context.Context, k Key[T]) (T, bool) {
	v, ok := ctx.Value(k).(T)
	return v, ok
}
