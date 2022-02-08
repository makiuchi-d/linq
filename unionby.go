package linq

type unionByEnumerator[T any, K comparable] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	kmap *keyMap[K, T]
}

// UnionBy produces the set union of two sequences according to a specified key selector function.
func UnionBy[T any, K comparable](first, second Enumerator[T], keySelector func(v T) (K, error)) Enumerator[T] {
	return &unionByEnumerator[T, K]{
		fst:  first,
		snd:  second,
		kmap: newKeyMap(keySelector),
	}
}

func (e *unionByEnumerator[T, K]) Next() (def T, _ error) {
	if e.fst != nil {
		v, err := e.fst.Next()
		if err == nil {
			if _, err = e.kmap.add(v); err != nil {
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
		has, err := e.kmap.has(v)
		if err != nil {
			return def, err
		}
		if !has {
			return v, nil
		}
	}
}
