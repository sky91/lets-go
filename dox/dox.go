package dox

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/samber/do/v2"
)

type Transformer[In, Out any] func(in In) (Out, error)
type Producer[Out any] func() (Out, error)

type componentKeyHelper struct{}
type ComponentKey *componentKeyHelper

type Ref[K ~*componentKeyHelper, V any] struct{ Val V }

func LazyWithRef[K ~*componentKeyHelper, V any](provider do.Provider[V]) func(do.Injector) {
	var buf [4]byte
	_, _ = rand.Read(buf[:])
	name := hex.EncodeToString(buf[:])
	return do.Package(
		do.LazyNamed(name, provider),
		do.Lazy(func(injector do.Injector) (Ref[K, V], error) {
			return Ref[K, V]{Val: do.MustInvokeNamed[V](injector, name)}, nil
		}),
	)
}

func Produce[T any](produce Producer[T]) do.Provider[T] {
	return func(injector do.Injector) (T, error) { return produce() }
}

func Transform[In, Out any](transform Transformer[In, Out]) do.Provider[Out] {
	return func(injector do.Injector) (out Out, err error) {
		in, err := do.Invoke[In](injector)
		if err != nil {
			return out, err
		}
		return transform(in)
	}
}

func StructTransform[In, Out any](transform Transformer[In, Out]) do.Provider[Out] {
	return func(injector do.Injector) (out Out, err error) {
		in, err := do.InvokeStruct[In](injector)
		if err != nil {
			return out, err
		}
		return transform(in)
	}
}

func StructSelf[T any]() do.Provider[T] {
	return StructTransform(func(in T) (T, error) { return in, nil })
}

func StructSelfAndInit[T, InitParam any](init func(t T, param InitParam) error) do.Provider[T] {
	return func(injector do.Injector) (T, error) {
		t, err := do.InvokeStruct[T](injector)
		if err != nil {
			return t, err
		}
		param, err := do.InvokeStruct[InitParam](injector)
		if err != nil {
			return t, err
		}
		return t, init(t, param)
	}
}
