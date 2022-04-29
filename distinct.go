package linq

type distinctEnumerator[T any, K comparable] struct {
	src  Enumerator[T]
	hmap *hashMap[K, T]
}

// Distinct returns distinct elements from a sequence by using the specified comparer functions.
func Distinct[T any, E IEnumerable[T]](src E, equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &distinctEnumerator[T, int]{
			src:  src(),
			hmap: newHashMap(getHashCode, equals),
		}
	}
}

// DistinctBy returns distinct elements from a sequence according to a specified key selector function.
func DistinctBy[T any, K comparable, E IEnumerable[T]](src E, keySelector func(v T) (K, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &distinctEnumerator[T, K]{
			src:  src(),
			hmap: newKeyMap(keySelector),
		}
	}
}

func (e *distinctEnumerator[T, int]) Next() (def T, _ error) {
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
