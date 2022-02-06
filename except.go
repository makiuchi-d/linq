package linq

type exceptEnumerator[T any] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	eq   func(T, T) (bool, error)
	hash func(T) (int, error)
	hmap *hashMap[int, T]
}

// Except produces the set difference of two sequences by using the specified comparer functions.
func Except[T any](first, second Enumerator[T], equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerator[T] {
	return &exceptEnumerator[T]{
		fst:  first,
		snd:  second,
		eq:   equals,
		hash: getHashCode,
	}
}

func (e *exceptEnumerator[T]) Next() (def T, _ error) {
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
		if !has {
			return v, nil
		}
	}
}
