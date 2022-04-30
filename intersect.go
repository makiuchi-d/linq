package linq

type intersectEnumerator[T any, H comparable] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	eq   func(T, T) (bool, error)
	hash func(T) (H, error)
	hmap *hashMap[H, T]
}

// Intersect produces the set intersection of two sequences by using the specified comparer functions.
func Intersect[T any, E IEnumerable[T]](first, second E, equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &intersectEnumerator[T, int]{
			fst:  first(),
			snd:  second(),
			eq:   equals,
			hash: getHashCode,
		}
	}
}

// IntersectBy produces the set intersection of two sequences according to a specified key selector function.
func IntersectBy[T any, K comparable, E IEnumerable[T]](first, second E, keySelector func(v T) (K, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &intersectEnumerator[T, K]{
			fst:  first(),
			snd:  second(),
			eq:   alwaysEqual[T],
			hash: keySelector,
		}
	}
}

func (e *intersectEnumerator[T, H]) Next() (def T, _ error) {
	if e.hmap == nil {
		hm := newHashMap(e.hash, e.eq)
		if err := hm.addAll(e.snd); err != nil {
			return def, err
		}
		e.hmap = hm
	}

	for {
		v, err := e.fst.Next()
		if err != nil {
			return def, err
		}
		has, err := e.hmap.has(v)
		if err != nil {
			return def, err
		}
		if has {
			return v, nil
		}
	}
}
