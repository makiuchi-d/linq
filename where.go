package linq

type whereEnumerator[T any] struct {
	src  Enumerator[T]
	pred func(v T) (bool, error)
}

// Where filters a sequence of values based on a predicate.
func Where[T any](src Enumerator[T], pred func(v T) (bool, error)) Enumerator[T] {
	return &whereEnumerator[T]{src: src, pred: pred}
}

func (e *whereEnumerator[T]) Next() (T, error) {
	for {
		v, err := e.src.Next()
		if err != nil {
			return v, err
		}

		ok, err := e.pred(v)
		if err != nil {
			var v T
			return v, err
		}
		if ok {
			return v, nil
		}
	}
}
