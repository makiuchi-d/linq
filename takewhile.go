package linq

type takeWhileEnumerator[T any] struct {
	src  Enumerator[T]
	pred func(T) (bool, error)
}

// TakeWhile returns elements from a sequence as long as a specified condition is true, and then skips the remaining elements.
func TakeWhile[T any](src Enumerator[T], pred func(T) (bool, error)) Enumerator[T] {
	return &takeWhileEnumerator[T]{src: src, pred: pred}
}

func (e *takeWhileEnumerator[T]) Next() (def T, _ error) {
	v, err := e.src.Next()
	if err != nil {
		return def, err
	}
	ok, err := e.pred(v)
	if err != nil {
		return def, err
	}
	if !ok {
		return def, EOC
	}
	return v, nil
}
