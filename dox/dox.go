package dox

import "github.com/samber/do/v2"

func LazyTransfer[Dep, T any](transfer func(dep Dep) (T, error)) func(do.Injector) {
	return func(injector do.Injector) {
		do.Provide(injector, func(injector do.Injector) (t T, err error) {
			dep, err := do.Invoke[Dep](injector)
			if err != nil {
				return t, err
			}
			return transfer(dep)
		})
	}
}

func LazyTransferNamed[Dep, T any](name string, transfer func(dep Dep) (T, error)) func(do.Injector) {
	return func(injector do.Injector) {
		do.ProvideNamed(injector, name, func(injector do.Injector) (t T, err error) {
			dep, err := do.Invoke[Dep](injector)
			if err != nil {
				return t, err
			}
			return transfer(dep)
		})
	}
}
