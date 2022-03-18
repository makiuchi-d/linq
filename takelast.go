package linq

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
