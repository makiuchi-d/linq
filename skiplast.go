package linq

type skipLastEnumerator[T any] struct {
	src Enumerator[T]
	cnt int
	buf []T
	i   int
}

// SkipLast returns a new enumerable collection that contains the elements from source with the last count elements of the source collection omitted.
func SkipLast[T any](src Enumerator[T], count int) Enumerator[T] {
	return &skipLastEnumerator[T]{src: src, cnt: count}
}

func (e *skipLastEnumerator[T]) Next() (def T, _ error) {
	if e.buf == nil {
		e.buf = make([]T, e.cnt)
		for i := 0; i < e.cnt; i++ {
			v, err := e.src.Next()
			if err != nil {
				return def, err
			}
			e.buf[i] = v
		}
	}

	i := e.i % e.cnt
	r := e.buf[i]
	v, err := e.src.Next()
	if err != nil {
		return def, err
	}
	e.buf[i] = v
	e.i++
	return r, nil
}
