package linq

type sliceEnumerator[T any] struct {
	s []T
	i int
}

// FromSlice generates an Enumerator[T] from a slice.
func FromSlice[S ~[]T, T any](s S) Enumerator[T] {
	return &sliceEnumerator[T]{s: []T(s)}
}

// ToSlice creates a slice from an Enumerator[T]
func ToSlice[T any](src Enumerator[T]) ([]T, error) {
	s := make([]T, 0)
	for {
		v, err := src.Next()
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
	v := e.s[e.i]
	e.i++
	return v, nil
}
