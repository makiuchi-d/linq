package linq

type skipWhileEnumerator[T any] struct {
	src     Enumerator[T]
	pred    func(T) (bool, error)
	skipped bool
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
func SkipWhile[T any](src Enumerator[T], pred func(T) (bool, error)) Enumerator[T] {
	return &skipWhileEnumerator[T]{src: src, pred: pred}
}

func (e *skipWhileEnumerator[T]) Next() (def T, _ error) {
	if e.skipped {
		return e.src.Next()
	}
	for {
		v, err := e.src.Next()
		if err != nil {
			return def, err
		}
		ok, err := e.pred(v)
		if err != nil {
			return def, err
		}
		if !ok {
			e.skipped = true
			return v, nil
		}
	}
}
