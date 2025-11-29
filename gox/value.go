package gox

type Value[T any] interface {
	Val() T
	Ptr() *T
}

func Nil2Zero[T any](t *T) T {
	if t == nil {
		var defaultValue T
		return defaultValue
	}
	return *t
}

func New[T any](t T) *T { return &t }

func DerefOr[T any](t *T, defaultVal T) T {
	if t == nil {
		return defaultVal
	}
	return *t
}

func DerefOrCompute[T any](t *T, compute func() T) T {
	if t == nil {
		return compute()
	}
	return *t
}
