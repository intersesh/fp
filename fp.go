package fp

func Unique[T comparable](items []T) []T {
	set := ToMap(items, func(t T) (T, struct{}) {
		return t, struct{}{}
	})
	return Keys(set)
}

func UniqueLinear[T any, K comparable](items []T, f func(T) K) []T {
	seen := make(map[K]struct{})
	out := make([]T, 0)
	for _, i := range items {
		k := f(i)
		if _, ok := seen[k]; !ok {
			out = append(out, i)
			seen[k] = struct{}{}
		}
	}

	return out
}

func ToMap[T any, K comparable, V any](slice []T, f func(T) (K, V)) map[K]V {
	m := make(map[K]V, len(slice))
	for _, l := range slice {
		k, v := f(l)
		m[k] = v
	}

	return m
}

func Keys[K comparable, V any](m map[K]V) []K {
	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}

	return s
}

func Values[K comparable, V any](m map[K]V) []V {
	s := make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}

	return s
}

func Map[T, U any](items []T, f func(T) U) []U {
	s := make([]U, 0, len(items))

	for _, i := range items {
		s = append(s, f(i))
	}

	return s
}

func Group[K comparable, V any](items []V, f func(V) K) map[K][]V {
	m := make(map[K][]V)

	for _, i := range items {
		k := f(i)
		m[k] = append(m[k], i)
	}

	return m
}

func All(items []bool) bool {
	for _, i := range items {
		if !i {
			return false
		}
	}

	return true
}

func Reduce[T, U any](items []T, f func(acc T, current U) U) U {
	var out U
	for _, i := range items {
		out = f(i, out)
	}

	return out
}
