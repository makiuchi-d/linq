package linq

type chunkEnumerator[T any] struct {
	src  Enumerator[T]
	size int
}

// Chunk splits the elements of a sequence into chunks of size at most `size`.
func Chunk[T any](src Enumerator[T], size int) Enumerator[[]T] {
	return &chunkEnumerator[T]{
		src:  src,
		size: size,
	}
}

func (e *chunkEnumerator[T]) Next() ([]T, error) {
	s := make([]T, 0, e.size)

	for i := 0; i < e.size; i++ {
		v, err := e.src.Next()
		if err != nil {
			if isEOC(err) {
				break
			}
			return nil, err
		}
		s = append(s, v)
	}

	if len(s) == 0 {
		return nil, EOC
	}

	return s, nil
}
