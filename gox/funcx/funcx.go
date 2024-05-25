package funcx

func DoIgnoreReturn1[T any](f func() T) { _ = f() }
