package lifecycle

import (
	"context"
	"errors"
)

type Closers []func()

func (thisP *Closers) Defer(f func()) {
	*thisP = append(*thisP, f)
}

func (thisP *Closers) Close() {
	for i := len(*thisP) - 1; i >= 0; i-- {
		(*thisP)[i]()
	}
	*thisP = nil
}

type Input1Closers[T any] []func(T)

func (thisP *Input1Closers[T]) Defer(f func(T)) {
	*thisP = append(*thisP, f)
}

func (thisP *Input1Closers[T]) Close(t T) {
	for i := len(*thisP) - 1; i >= 0; i-- {
		(*thisP)[i](t)
	}
	*thisP = nil
}

type ErrClosers []func() error

func (thisP *ErrClosers) Defer(f func()) {
	thisP.DeferErr(func() error { f(); return nil })
}

func (thisP *ErrClosers) DeferErr(f func() error) {
	*thisP = append(*thisP, f)
}

func (thisP *ErrClosers) Close() error {
	errs := make([]error, 0, len(*thisP))
	for i := len(*thisP) - 1; i >= 0; i-- {
		if e := (*thisP)[i](); e != nil {
			errs = append(errs, e)
		}
	}
	*thisP = nil
	return errors.Join(errs...)
}

func (thisP *ErrClosers) CloseOnError(err *error) error {
	if err == nil || *err == nil {
		return nil
	}
	return thisP.Close()
}

func (thisP *ErrClosers) CloseAppendOnError(err *error) {
	if err == nil || *err == nil {
		return
	}
	if closeErr := thisP.Close(); closeErr != nil {
		*err = errors.Join(*err, closeErr)
	}
}

type Input1ErrClosers[T any] []func(T) error

func (thisP *Input1ErrClosers[T]) Defer(f func(T)) {
	thisP.DeferErr(func(t T) error { f(t); return nil })
}

func (thisP *Input1ErrClosers[T]) DeferErr(f func(T) error) {
	*thisP = append(*thisP, f)
}

func (thisP *Input1ErrClosers[T]) Close(t T) error {
	errs := make([]error, 0, len(*thisP))
	for i := len(*thisP) - 1; i >= 0; i-- {
		if e := (*thisP)[i](t); e != nil {
			errs = append(errs, e)
		}
	}
	*thisP = nil
	return errors.Join(errs...)
}

func (thisP *Input1ErrClosers[T]) CloseOnError(t T, err *error) error {
	if err == nil || *err == nil {
		return nil
	}
	return thisP.Close(t)
}

func (thisP *Input1ErrClosers[T]) CloseAppendOnError(t T, err *error) {
	if err == nil || *err == nil {
		return
	}
	if closeErr := thisP.Close(t); closeErr != nil {
		*err = errors.Join(*err, closeErr)
	}
}

type CtxClosers = Input1Closers[context.Context]

type CtxErrClosers = Input1ErrClosers[context.Context]
