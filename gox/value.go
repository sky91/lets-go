package gox

type Value[T any] interface {
	Val() T
	Ptr() *T
}
