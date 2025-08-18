package fp

func Filter[T any](items []T, removers ...func(T) bool) []T {
	out := make([]T, 0, len(items))

itemLoop:
	for _, i := range items {
		for _, f := range removers {
			if f(i) {
				continue itemLoop
			}
		}

		out = append(out, i)
	}

	return out
}

func Map[T, U any](items []T, f func(T) U) []U {
	out := make([]U, 0, len(items))
	for _, i := range items {
		out = append(out, f(i))
	}

	return out
}
