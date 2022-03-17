package linq

type skipEnumerator[T any] struct {
	src Enumerator[T]
	cnt int
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
func Skip[T any](src Enumerator[T], count int) Enumerator[T] {
	return &skipEnumerator[T]{src: src, cnt: count}
}

func (e *skipEnumerator[T]) Next() (def T, _ error) {
	for ; e.cnt > 0; e.cnt-- {
		_, err := e.src.Next()
		if err != nil {
			return def, err
		}
	}
	return e.src.Next()
}
