package lifecycle

import "context"

type FnShutdown[T any] struct {
	Value      T
	ShutdownFn func()
}

func (thisP *FnShutdown[T]) Val() T    { return thisP.Value }
func (thisP *FnShutdown[T]) Ptr() *T   { return &thisP.Value }
func (thisP *FnShutdown[T]) Shutdown() { thisP.ShutdownFn() }

type ErrFnShutdown[T any] struct {
	Value      T
	ShutdownFn func() error
}

func (thisP *ErrFnShutdown[T]) Val() T          { return thisP.Value }
func (thisP *ErrFnShutdown[T]) Ptr() *T         { return &thisP.Value }
func (thisP *ErrFnShutdown[T]) Shutdown() error { return thisP.ShutdownFn() }

type CtxFnShutdown[T any] struct {
	Value      T
	ShutdownFn func(context.Context)
}

func (thisP *CtxFnShutdown[T]) Val() T                       { return thisP.Value }
func (thisP *CtxFnShutdown[T]) Ptr() *T                      { return &thisP.Value }
func (thisP *CtxFnShutdown[T]) Shutdown(ctx context.Context) { thisP.ShutdownFn(ctx) }

type CtxErrFnShutdown[T any] struct {
	Value      T
	ShutdownFn func(context.Context) error
}

func (thisP *CtxErrFnShutdown[T]) Val() T                             { return thisP.Value }
func (thisP *CtxErrFnShutdown[T]) Ptr() *T                            { return &thisP.Value }
func (thisP *CtxErrFnShutdown[T]) Shutdown(ctx context.Context) error { return thisP.ShutdownFn(ctx) }

type CloserShutdown[T interface{ Close() }] struct{ Value T }

func (thisV CloserShutdown[T]) Val() T    { return thisV.Value }
func (thisV CloserShutdown[T]) Ptr() *T   { return &thisV.Value }
func (thisV CloserShutdown[T]) Shutdown() { thisV.Value.Close() }

type ErrCloserShutdown[T interface{ Close() error }] struct{ Value T }

func (thisV ErrCloserShutdown[T]) Val() T          { return thisV.Value }
func (thisV ErrCloserShutdown[T]) Ptr() *T         { return &thisV.Value }
func (thisV ErrCloserShutdown[T]) Shutdown() error { return thisV.Value.Close() }

type CtxCloserShutdown[T interface{ Close(context.Context) }] struct{ Value T }

func (thisV CtxCloserShutdown[T]) Val() T                       { return thisV.Value }
func (thisV CtxCloserShutdown[T]) Ptr() *T                      { return &thisV.Value }
func (thisV CtxCloserShutdown[T]) Shutdown(ctx context.Context) { thisV.Value.Close(ctx) }

type CtxErrCloserShutdown[T interface{ Close(context.Context) error }] struct{ Value T }

func (thisV CtxErrCloserShutdown[T]) Val() T  { return thisV.Value }
func (thisV CtxErrCloserShutdown[T]) Ptr() *T { return &thisV.Value }
func (thisV CtxErrCloserShutdown[T]) Shutdown(ctx context.Context) error {
	return thisV.Value.Close(ctx)
}
