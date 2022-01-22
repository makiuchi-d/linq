package linq

type zipEnumerator[T, U, V any] struct {
	first  Enumerator[T]
	second Enumerator[U]
	sel    func(T, U) (V, error)
}

// Zip applies a specified function to the corresponding elements of two sequences, producing a sequence of the results.
func Zip[T, U, V any](first Enumerator[T], second Enumerator[U], resultSelector func(T, U) (V, error)) Enumerator[V] {
	return &zipEnumerator[T, U, V]{
		first:  first,
		second: second,
		sel:    resultSelector,
	}
}

func (e *zipEnumerator[T, U, V]) Next() (V, error) {
	var d V

	t, err := e.first.Next()
	if err != nil {
		return d, err
	}

	u, err := e.second.Next()
	if err != nil {
		return d, err
	}

	return e.sel(t, u)
}
