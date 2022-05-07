package linq

type selectManyEnumerator[S, C, T any] struct {
	src  Enumerator[S]
	csel func(S) (Enumerable[C], error)
	rsel func(C) (T, error)

	cur Enumerator[C]
}

// SelectMany projects each element of a sequence to an Enumerable[T] and flattens the resulting sequences into one sequence.
func SelectMany[S, C, T any, E IEnumerable[S], EC IEnumerable[C]](src E, collectionSelector func(S) (EC, error), resultSelector func(C) (T, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &selectManyEnumerator[S, C, T]{
			src: src(),
			csel: func(s S) (Enumerable[C], error) {
				c, err := collectionSelector(s)
				return Enumerable[C](c), err
			},
			rsel: resultSelector,
		}
	}
}

func (e *selectManyEnumerator[S, C, T]) Next() (def T, _ error) {
	if e.cur == nil {
		t, err := e.src.Next()
		if err != nil {
			return def, err // includes case of EndOfCollection
		}

		c, err := e.csel(t)
		if err != nil {
			return def, err
		}

		e.cur = c()
	}

	u, err := e.cur.Next()
	if err != nil {
		if isEOC(err) {
			e.cur = nil
			return e.Next()
		}
		return def, err
	}

	return e.rsel(u)
}
