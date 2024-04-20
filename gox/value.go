package gox

type Value[T any] interface {
	Val() T
	Ptr() *T
}

func Nil2Zero[T any](t *T) (r T) {
	if t == nil {
		return
	}
	return *t
}
