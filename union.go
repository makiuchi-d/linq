package linq

type unionEnumerator[T any] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	hmap *hashMap[int, T]
}

// Union produces the set union of two sequences by using the specified comparer functions.
func Union[T any](first, second Enumerator[T], equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerator[T] {
	return &unionEnumerator[T]{
		fst:  first,
		snd:  second,
		hmap: newHashMap(getHashCode, equals),
	}
}

func (e *unionEnumerator[T]) Next() (def T, _ error) {
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
