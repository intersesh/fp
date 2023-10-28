package fp

type F[T any] func(T) T

func Compose[T any](fl ...F[T]) F[T] {
	return func(x T) (t T) {
		result := x

		for _, f := range fl {
			result = f(result)
		}

		return result
	}
}

func Combine[T, U, V any](f1 func(T) U, f2 func(U) V) func(T) V {
	return func(t T) (v V) {
		return f2(f1(t))
	}
}

type FE[T any] func(T) (T, error)

func ComposeError[T any](fl ...FE[T]) FE[T] {
	return func(x T) (t T, err error) {
		result := x

		for _, f := range fl {
			result, err = f(result)
			if err != nil {
				return t, err
			}
		}

		return result, nil
	}
}

func CombineError[T, U, V any](
	f1 func(T) (U, error),
	f2 func(U) (V, error),
) func(T) (V, error) {
	return func(t T) (v V, err error) {
		var u U
		u, err = f1(t)
		if err != nil {
			return v, err
		}

		return f2(u)
	}
}
