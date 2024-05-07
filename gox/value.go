package gox

type Value[T any] interface {
	Val() T
	Ptr() *T
}

func Nil2Zero[T any](t *T) T {
	if t == nil {
		return *new(T)
	}
	return *t
}

func New[T any](t T) *T { return &t }
