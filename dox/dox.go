package dox

import (
	"github.com/samber/do/v2"
)

type Transformer[In, Out any] func(in In) (Out, error)
type Producer[Out any] func() (Out, error)

func LazyProvide[T any](provider do.Provider[T], names ...string) func(do.Injector) {
	if len(names) == 0 {
		return func(injector do.Injector) { do.Provide(injector, provider) }
	}
	return func(injector do.Injector) {
		for _, name := range names {
			do.ProvideNamed(injector, name, provider)
		}
	}
}

func LazyProduce[Out any](produce Producer[Out], names ...string) func(do.Injector) {
	return LazyProvide(func(do.Injector) (out Out, err error) {
		return produce()
	}, names...)
}

func LazyTransform[In, Out any](transform Transformer[In, Out], names ...string) func(do.Injector) {
	return LazyProvide(func(injector do.Injector) (out Out, err error) {
		in, err := do.Invoke[In](injector)
		if err != nil {
			return out, err
		}
		return transform(in)
	}, names...)
}

func LazyStructTransform[In, Out any](transform Transformer[In, Out], names ...string) func(do.Injector) {
	return LazyProvide(func(injector do.Injector) (out Out, err error) {
		in, err := do.InvokeStruct[In](injector)
		if err != nil {
			return out, err
		}
		return transform(in)
	}, names...)
}

func LazyStructSelf[T any](names ...string) func(do.Injector) {
	return LazyStructTransform(func(in T) (T, error) { return in, nil }, names...)
}

func LazyStructSelfAndInit[T, InitParam any](init func(t T, param InitParam) error, names ...string) func(do.Injector) {
	return LazyProvide(func(injector do.Injector) (T, error) {
		t, err := do.InvokeStruct[T](injector)
		if err != nil {
			return t, err
		}
		param, err := do.InvokeStruct[InitParam](injector)
		if err != nil {
			return t, err
		}
		return t, init(t, param)
	}, names...)
}

func LazyUnwrap[In wrapper[Out], Out any](names ...string) func(do.Injector) {
	return LazyProvide(func(injector do.Injector) (out Out, err error) {
		w, err := do.Invoke[In](injector)
		if err != nil {
			return out, err
		}
		return w.Val(), nil
	}, names...)
}

func LazyWithUnwrap[In, Out any, W wrapper[Out]](transform Transformer[In, W], names ...string) func(do.Injector) {
	return do.Package(LazyTransform(transform, names...), LazyUnwrap[W, Out](names...))
}

func LazyStructWithUnwrap[In, Out any, W wrapper[Out]](transform Transformer[In, W], names ...string) func(do.Injector) {
	return do.Package(LazyStructTransform(transform, names...), LazyUnwrap[W, Out](names...))
}

type wrapper[T any] interface {
	Val() T
}
