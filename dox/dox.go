package dox

import (
	"github.com/samber/do/v2"
)

type Transformer[In, Out any] func(in In) (Out, error)
type Producer[Out any] func() (Out, error)

func LazyProduce[Out any](produce Producer[Out]) func(do.Injector) {
	return func(injector do.Injector) {
		do.Provide(injector, func(do.Injector) (out Out, err error) {
			return produce()
		})
	}
}

func Lazy[In, Out any](transform Transformer[In, Out]) func(do.Injector) {
	return func(injector do.Injector) {
		do.Provide(injector, func(injector do.Injector) (out Out, err error) {
			in, err := do.Invoke[In](injector)
			if err != nil {
				return out, err
			}
			return transform(in)
		})
	}
}

func LazyNamed[In, Out any](name string, transform Transformer[In, Out]) func(do.Injector) {
	return func(injector do.Injector) {
		do.ProvideNamed(injector, name, func(injector do.Injector) (out Out, err error) {
			in, err := do.Invoke[In](injector)
			if err != nil {
				return out, err
			}
			return transform(in)
		})
	}
}

func LazyStruct[In, Out any](transform Transformer[*In, Out]) func(do.Injector) {
	return func(injector do.Injector) {
		do.Provide(injector, func(injector do.Injector) (out Out, err error) {
			in, err := do.InvokeStruct[In](injector)
			if err != nil {
				return out, err
			}
			return transform(in)
		})
	}
}

func LazyStructSelf[T any]() func(do.Injector) {
	return LazyStruct[T, *T](TransformSelf[*T])
}

func LazyStructNamed[In, Out any](name string, transform Transformer[*In, Out]) func(do.Injector) {
	return func(injector do.Injector) {
		do.ProvideNamed(injector, name, func(injector do.Injector) (out Out, err error) {
			in, err := do.InvokeStruct[In](injector)
			if err != nil {
				return out, err
			}
			return transform(in)
		})
	}
}

func LazyUnwrap[In wrapper[Out], Out any]() func(do.Injector) {
	return func(injector do.Injector) {
		do.Provide(injector, func(injector do.Injector) (out Out, err error) {
			w, err := do.Invoke[In](injector)
			if err != nil {
				return out, err
			}
			return w.Val(), nil
		})
	}
}

func LazyWithUnwrap[In, Out any, W wrapper[Out]](transform Transformer[In, W]) func(do.Injector) {
	return do.Package(Lazy(transform), LazyUnwrap[W, Out]())
}

func LazyStructWithUnwrap[In, Out any, W wrapper[Out]](transform Transformer[*In, W]) func(do.Injector) {
	return do.Package(LazyStruct(transform), LazyUnwrap[W, Out]())
}

type wrapper[T any] interface {
	Val() T
}

func TransformSelf[T any](t T) (T, error) {
	return t, nil
}
