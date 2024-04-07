package gox

type Value[T any] interface{ Get() T }
