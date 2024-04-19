package panicx

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[T any](t T, err error) T {
	Must(err)
	return t
}

func Must2[T, R any](t T, r R, err error) (T, R) {
	Must(err)
	return t, r
}
