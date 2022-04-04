package linq

type exceptEnumerator[T any, H comparable] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	eq   func(T, T) (bool, error)
	hash func(T) (H, error)
	hmap *hashMap[H, T]
}

// Except produces the set difference of two sequences by using the specified comparer functions.
func Except[T any](first, second Enumerator[T], equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerator[T] {
	return &exceptEnumerator[T, int]{
		fst:  first,
		snd:  second,
		eq:   equals,
		hash: getHashCode,
	}
}

// ExceptBy produces the set difference of two sequences according to a specified key selector function.
func ExceptBy[T any, K comparable](first, second Enumerator[T], keySelector func(v T) (K, error)) Enumerator[T] {
	return &exceptEnumerator[T, K]{
		fst:  first,
		snd:  second,
		eq:   alwaysEqual[T],
		hash: keySelector,
	}
}

func (e *exceptEnumerator[T, H]) Next() (def T, _ error) {
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
