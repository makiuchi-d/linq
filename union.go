package linq

type unionEnumerator[T any, H comparable] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	hmap *hashMap[H, T]
}

// Union produces the set union of two sequences by using the specified comparer functions.
func Union[T any, E IEnumerable[T]](first, second E, equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &unionEnumerator[T, int]{
			fst:  first(),
			snd:  second(),
			hmap: newHashMap(getHashCode, equals),
		}
	}
}

// UnionBy produces the set union of two sequences according to a specified key selector function.
func UnionBy[T any, K comparable, E IEnumerable[T]](first, second E, keySelector func(v T) (K, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &unionEnumerator[T, K]{
			fst:  first(),
			snd:  second(),
			hmap: newKeyMap(keySelector),
		}
	}
}

func (e *unionEnumerator[T, H]) Next() (def T, _ error) {
	if e.fst != nil {
		v, err := e.fst.Next()
		if err == nil {
			if _, err = e.hmap.add(v); err != nil {
				return def, err
			}
			return v, nil
		}
		if !isEOC(err) {
			return def, err
		}
		e.fst = nil
	}

	for {
		v, err := e.snd.Next()
		if err != nil {
			return def, err
		}
		has, err := e.hmap.has(v)
		if err != nil {
			return def, err
		}
		if !has {
			return v, nil
		}
	}
}
