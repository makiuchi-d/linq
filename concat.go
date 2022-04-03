package linq

type concatEnumerator[T any] struct {
	fst Enumerator[T]
	snd Enumerator[T]
}

// Concat concatenates two sequences.
func Concat[T any](first, second Enumerator[T]) Enumerator[T] {
	return &concatEnumerator[T]{fst: first, snd: second}
}

func (e *concatEnumerator[T]) Next() (def T, _ error) {
	v, err := e.fst.Next()
	if err == nil {
		return v, nil
	}
	if !isEOC(err) {
		return def, err
	}
	if e.snd == nil {
		return def, EOC
	}
	e.fst, e.snd = e.snd, nil
	return e.Next()
}
