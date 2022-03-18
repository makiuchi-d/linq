package linq

type takeEnumerator[T any] struct {
	src Enumerator[T]
	cnt int
}

// Take returns a specified number of contiguous elements from the start of a sequence.
func Take[T any](src Enumerator[T], count int) Enumerator[T] {
	return &takeEnumerator[T]{src: src, cnt: count}
}

func (e *takeEnumerator[T]) Next() (def T, _ error) {
	if e.cnt <= 0 {
		return def, EOC
	}
	e.cnt--
	return e.src.Next()
}
