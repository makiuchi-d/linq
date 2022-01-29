package linq

type distinctByEnumerator[T any, K comparable] struct {
	src    Enumerator[T]
	keysel func(T) (K, error)
	keymap map[K]struct{}
}

// DistinctBy returns distinct elements from a sequence according to a specified key selector function.
func DistinctBy[T any, K comparable](src Enumerator[T], keySelector func(v T) (K, error)) Enumerator[T] {
	return &distinctByEnumerator[T, K]{
		src:    src,
		keysel: keySelector,
		keymap: make(map[K]struct{}),
	}
}

func (e *distinctByEnumerator[T, K]) Next() (T, error) {
	for {
		v, err := e.src.Next()
		if err != nil {
			return v, err
		}

		k, err := e.keysel(v)
		if err != nil {
			var d T
			return d, err
		}

		_, ok := e.keymap[k]
		if !ok {
			e.keymap[k] = struct{}{}
			return v, nil
		}
	}
}
