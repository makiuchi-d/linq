package linq

type intersectByEnumerator[T any, K comparable] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	ksel func(T) (K, error)
	kmap *keyMap[K, T]
}

// IntersectBy produces the set intersection of two sequences according to a specified key selector function.
func IntersectBy[T any, K comparable](first, second Enumerator[T], keySelector func(v T) (K, error)) Enumerator[T] {
	return &intersectByEnumerator[T, K]{
		fst:  first,
		snd:  second,
		ksel: keySelector,
	}
}

func (e *intersectByEnumerator[T, K]) Next() (def T, _ error) {
	if e.kmap == nil {
		km := newKeyMap(e.ksel)
		if err := km.addAll(e.snd); err != nil {
			return def, err
		}
		e.kmap = km
	}

	for {
		v, err := e.fst.Next()
		if err != nil {
			return def, err
		}
		has, err := e.kmap.has(v)
		if err != nil {
			return def, err
		}
		if has {
			return v, nil
		}
	}
}
