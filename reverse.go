package linq

type reverseEnumerator[T any] struct {
	src Enumerator[T]
	s   []T
	i   int
}

// Reverse inverts the order of the elements in a sequence.
func Reverse[T any, E IEnumerable[T]](src E) Enumerable[T] {
	return func() Enumerator[T] {
		return &reverseEnumerator[T]{
			src: src(),
		}
	}
}

func (e *reverseEnumerator[T]) Next() (def T, _ error) {
	if e.s == nil {
		s, err := toSlice(e.src)
		if err != nil {
			return def, err
		}
		e.s = s
	}

	l := len(e.s)
	i := e.i
	if i >= l {
		return def, EOC
	}
	e.i++
	return e.s[l-i-1], nil
}
