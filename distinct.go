package linq

type distinctEnumerator[T any] struct {
	src     Enumerator[T]
	equals  func(T, T) (bool, error)
	hash    func(T) (int, error)
	hashMap map[int]T
}

// Distinct returns distinct elements from a sequence by using the specified comparer functions.
func Distinct[T any](src Enumerator[T], equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerator[T] {
	return &distinctEnumerator[T]{
		src:     src,
		equals:  equals,
		hash:    getHashCode,
		hashMap: make(map[int]T),
	}
}

func (e *distinctEnumerator[T]) Next() (T, error) {
	for {
		v, err := e.src.Next()
		if err != nil {
			var d T
			return d, err
		}

		h, err := e.hash(v)
		if err != nil {
			var d T
			return d, err
		}

		t, ok := e.hashMap[h]
		if !ok {
			e.hashMap[h] = v
			return v, nil
		}

		eq, err := e.equals(t, v)
		if err != nil {
			var d T
			return d, nil
		}
		if !eq {
			return v, nil
		}
	}
}
