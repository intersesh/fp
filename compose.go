package fp

func Compose[T any](fl ...T) func(T) T {
	return func(x T) (t T) {
		result := x

		for _, f := range fl {
			result = f(result)
		}

		return result
	}
}

func ComposeError[T any](fl ...func(T) (T, error)) func(T) (T, error) {
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

func Combine2[T, U, V any](
	f1 func(T) U,
	f2 func(U) V,
) func(T) V {
	return func(t T) (v V) {
		return f2(f1(t))
	}
}

func Combine3[A, B, C, D any](
	f1 func(A) B,
	f2 func(B) C,
	f3 func(C) D,
) func(A) D {
	return func(a A) D {
		return f3(f2(f1(a)))
	}
}

func Combine4[A, B, C, D, E any](
	f1 func(A) B,
	f2 func(B) C,
	f3 func(C) D,
	f4 func(D) E,
) func(A) E {
	return func(a A) E {
		return f4(f3(f2(f1(a))))
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

func CombineError3[A, B, C, D any](
	f1 func(A) (B, error),
	f2 func(B) (C, error),
	f3 func(C) (D, error),
) func(A) (D, error) {
	return func(a A) (d D, err error) {
		var (
			b B
			c C
		)

		b, err = f1(a)
		if err != nil {
			return d, err
		}

		c, err = f2(b)
		if err != nil {
			return d, err
		}

		return f3(c)
	}
}

func CombineError4[A, B, C, D, E any](
	f1 func(A) (B, error),
	f2 func(B) (C, error),
	f3 func(C) (D, error),
	f4 func(D) (E, error),
) func(A) (E, error) {
	return func(a A) (e E, err error) {
		var (
			b B
			c C
			d D
		)

		b, err = f1(a)
		if err != nil {
			return e, err
		}

		c, err = f2(b)
		if err != nil {
			return e, err
		}

		d, err = f3(c)
		if err != nil {
			return e, err
		}

		return f4(d)
	}
}
