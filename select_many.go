package linq

type selectManyEnumerator[T, U, V any] struct {
	src  Enumerator[T]
	csel func(T) (Enumerator[U], error)
	rsel func(U) (V, error)

	cur Enumerator[U]
}

// SelectMany projects each element of a sequence to an Enumerable[V] and flattens the resulting sequences into one sequence.
func SelectMany[T, U, V any](src Enumerator[T], collectionSelector func(T) (Enumerator[U], error), resultSelector func(U) (V, error)) Enumerator[V] {
	return &selectManyEnumerator[T, U, V]{
		src:  src,
		csel: collectionSelector,
		rsel: resultSelector,
	}
}

func (e *selectManyEnumerator[T, U, V]) Next() (V, error) {
	if e.cur == nil {
		t, err := e.src.Next()
		if err != nil {
			var d V
			return d, err // includes case of EndOfCollection
		}

		c, err := e.csel(t)
		if err != nil {
			var d V
			return d, err
		}

		e.cur = c
	}

	u, err := e.cur.Next()
	if err != nil {
		if isEOC(err) {
			e.cur = nil
			return e.Next()
		}
		var d V
		return d, err
	}

	return e.rsel(u)
}
