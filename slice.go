package linq

type sliceEnumerator[T any] struct {
	s []T
	i int
}

// FromSlice generates an IEnumerable[T] from a slice.
func FromSlice[S ~[]T, T any](s S) Enumerable[T] {
	return func() Enumerator[T] {
		return &sliceEnumerator[T]{s: []T(s)}
	}
}

// ToSlice creates a slice from an IEnumerable[T]
func ToSlice[T any, E IEnumerable[T]](src E) ([]T, error) {
	return toSlice(src())
}

func toSlice[T any](e Enumerator[T]) ([]T, error) {
	s := make([]T, 0)
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				break
			}
			return s, err
		}
		s = append(s, v)
	}
	return s, nil
}

func (e *sliceEnumerator[T]) Next() (def T, _ error) {
	if e.i >= len(e.s) {
		return def, EOC
	}
	i := e.i
	e.i++
	return e.s[i], nil
}
