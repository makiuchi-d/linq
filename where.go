package linq

type whereEnumerator[T any] struct {
	src  Enumerator[T]
	pred func(v T) (bool, error)
}

// Where filters a sequence of values based on a predicate.
func Where[T any, E IEnumerable[T]](src E, pred func(v T) (bool, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &whereEnumerator[T]{src: src(), pred: pred}
	}
}

func (e *whereEnumerator[T]) Next() (def T, _ error) {
	for {
		v, err := e.src.Next()
		if err != nil {
			return def, err
		}

		ok, err := e.pred(v)
		if err != nil {
			return def, err
		}
		if ok {
			return v, nil
		}
	}
}
