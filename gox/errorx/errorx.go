package errorx

import "github.com/sky91/lets-go/gox/fnx"

func DoIgnoreError(f func() error) { fnx.DoIgnoreReturn1(f) }

func OrZero[T any](t T, err error) T {
	if err == nil {
		return t
	}
	return *new(T)
}
