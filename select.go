package linq

type selectEnumerator[S, T any] struct {
	src Enumerator[S]
	sel func(v S) (T, error)
}

// Select projects each element of a sequence into a new form.
func Select[S, T any, E IEnumerable[S]](src E, selector func(v S) (T, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &selectEnumerator[S, T]{src: src(), sel: selector}
	}
}

func (e *selectEnumerator[S, T]) Next() (def T, _ error) {
	v, err := e.src.Next()
	if err != nil {
		return def, err
	}

	return e.sel(v)
}
