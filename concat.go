package linq

type concatEnumerator[T any] struct {
	fst Enumerator[T]
	snd Enumerable[T]
}

// Concat concatenates two sequences.
func Concat[T any, E IEnumerable[T]](first, second E) Enumerable[T] {
	return func() Enumerator[T] {
		return &concatEnumerator[T]{fst: first(), snd: Enumerable[T](second)}
	}
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
	e.fst, e.snd = e.snd(), nil
	return e.Next()
}
