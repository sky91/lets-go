package lifecycle

import "context"

type FnShutdown[T any] struct {
	Value      T
	ShutdownFn func()
}

func (thisP *FnShutdown[T]) Get() T { return thisP.Value }

func (thisP *FnShutdown[T]) Shutdown() { thisP.ShutdownFn() }

type ErrFnShutdown[T any] struct {
	Value      T
	ShutdownFn func() error
}

func (thisP *ErrFnShutdown[T]) Get() T { return thisP.Value }

func (thisP *ErrFnShutdown[T]) Shutdown() error { return thisP.ShutdownFn() }

type CtxFnShutdown[T any] struct {
	Value      T
	ShutdownFn func(context.Context)
}

func (thisP *CtxFnShutdown[T]) Get() T { return thisP.Value }

func (thisP *CtxFnShutdown[T]) Shutdown(ctx context.Context) { thisP.ShutdownFn(ctx) }

type CtxErrFnShutdown[T any] struct {
	Value      T
	ShutdownFn func(context.Context) error
}

func (thisP *CtxErrFnShutdown[T]) Get() T { return thisP.Value }

func (thisP *CtxErrFnShutdown[T]) Shutdown(ctx context.Context) error { return thisP.ShutdownFn(ctx) }

type CloserShutdown[T interface{ Close() }] struct{ Value T }

func (thisP *CloserShutdown[T]) Get() T { return thisP.Value }

func (thisP *CloserShutdown[T]) Shutdown() { thisP.Value.Close() }

type ErrCloserShutdown[T interface{ Close() error }] struct{ Value T }

func (thisP *ErrCloserShutdown[T]) Get() T { return thisP.Value }

func (thisP *ErrCloserShutdown[T]) Shutdown() error { return thisP.Value.Close() }

type CtxCloserShutdown[T interface{ Close(context.Context) }] struct{ Value T }

func (thisP *CtxCloserShutdown[T]) Get() T { return thisP.Value }

func (thisP *CtxCloserShutdown[T]) Shutdown(ctx context.Context) { thisP.Value.Close(ctx) }

type CtxErrCloserShutdown[T interface{ Close(context.Context) error }] struct{ Value T }

func (thisP *CtxErrCloserShutdown[T]) Get() T { return thisP.Value }

func (thisP *CtxErrCloserShutdown[T]) Shutdown(ctx context.Context) error {
	return thisP.Value.Close(ctx)
}
