package linq

type intersectEnumerator[T any] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	eq   func(T, T) (bool, error)
	hash func(T) (int, error)
	hmap *hashMap[int, T]
}

// Intersect produces the set intersection of two sequences by using the specified comparer functions.
func Intersect[T any](first, second Enumerator[T], equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerator[T] {
	return &intersectEnumerator[T]{
		fst:  first,
		snd:  second,
		eq:   equals,
		hash: getHashCode,
	}
}

func (e *intersectEnumerator[T]) Next() (def T, _ error) {
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
