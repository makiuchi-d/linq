package linq

type zipEnumerator[S1, S2, T any] struct {
	first  Enumerator[S1]
	second Enumerator[S2]
	sel    func(S1, S2) (T, error)
}

// Zip applies a specified function to the corresponding elements of two sequences, producing a sequence of the results.
func Zip[S1, S2, T any, E1 IEnumerable[S1], E2 IEnumerable[S2]](first E1, second E2, resultSelector func(S1, S2) (T, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &zipEnumerator[S1, S2, T]{
			first:  first(),
			second: second(),
			sel:    resultSelector,
		}
	}
}

func (e *zipEnumerator[S1, S2, T]) Next() (def T, _ error) {
	t, err := e.first.Next()
	if err != nil {
		return def, err
	}

	u, err := e.second.Next()
	if err != nil {
		return def, err
	}

	return e.sel(t, u)
}
