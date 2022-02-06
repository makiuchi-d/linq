package linq

type distinctEnumerator[T any] struct {
	src  Enumerator[T]
	hmap *hashMap[int, T]
}

// Distinct returns distinct elements from a sequence by using the specified comparer functions.
func Distinct[T any](src Enumerator[T], equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerator[T] {
	return &distinctEnumerator[T]{
		src:  src,
		hmap: newHashMap(getHashCode, equals),
	}
}

func (e *distinctEnumerator[T]) Next() (def T, _ error) {
	for {
		v, err := e.src.Next()
		if err != nil {
			return def, err
		}

		added, err := e.hmap.add(v)
		if err != nil {
			return def, err
		}
		if added {
			return v, nil
		}
	}
}
