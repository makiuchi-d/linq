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

type takeLastEnumerator[T any] struct {
	src Enumerator[T]
	cnt int
	buf []T
	ofs int
	i   int
}

// TakeLast returns a new enumerable collection that contains the last count elements from source.
func TakeLast[T any](src Enumerator[T], count int) Enumerator[T] {
	return &takeLastEnumerator[T]{src: src, cnt: count}
}

func (e *takeLastEnumerator[T]) Next() (def T, _ error) {
	if e.buf == nil {
		e.buf = make([]T, e.cnt)
		i := 0
		for ; ; i++ {
			v, err := e.src.Next()
			if err != nil {
				if isEOC(err) {
					break
				}
				return def, err
			}
			e.buf[i%e.cnt] = v
		}
		if i < e.cnt {
			e.buf = e.buf[:i]
		} else {
			e.ofs = i
		}
	}
	i := e.i
	if i >= len(e.buf) {
		return def, EOC
	}
	e.i++
	return e.buf[(e.ofs+i)%len(e.buf)], nil
}
