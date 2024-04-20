package errorx

func DoIgnoreError(f func() error) { _ = f() }

func OrZero[T any](t T, err error) T {
	if err == nil {
		return t
	}
	return *new(T)
}
