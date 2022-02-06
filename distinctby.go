package linq

type distinctByEnumerator[T any, K comparable] struct {
	src  Enumerator[T]
	kmap *keyMap[K, T]
}

// DistinctBy returns distinct elements from a sequence according to a specified key selector function.
func DistinctBy[T any, K comparable](src Enumerator[T], keySelector func(v T) (K, error)) Enumerator[T] {
	return &distinctByEnumerator[T, K]{
		src:  src,
		kmap: newKeyMap(keySelector),
	}
}

func (e *distinctByEnumerator[T, K]) Next() (def T, _ error) {
	for {
		v, err := e.src.Next()
		if err != nil {
			return def, err
		}

		added, err := e.kmap.add(v)
		if err != nil {
			return def, err
		}
		if added {
			return v, nil
		}
	}
}
