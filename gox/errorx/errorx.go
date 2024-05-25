package errorx

import "github.com/sky91/lets-go/gox/funcx"

func DoIgnoreError(f func() error) { funcx.DoIgnoreReturn1(f) }

func OrZero[T any](t T, err error) T {
	if err == nil {
		return t
	}
	return *new(T)
}
