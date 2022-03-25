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
