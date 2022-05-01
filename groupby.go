package linq

// Grouping represents a collection of objects that have a common key.
type Grouping[T any, K comparable] struct {
	Enumerable Enumerable[T]
	Key        K
}

type groupByEnumerator[T any, K comparable] struct {
	src  Enumerable[T]
	ksel func(T) (K, error)

	ks []K
	m  map[K][]T
	i  int
}

// GroupBy groups the elements of a sequence according to a specified key selector function.
func GroupBy[T any, K comparable, E IEnumerable[T]](src E, keySelector func(T) (K, error)) Enumerable[Grouping[T, K]] {
	return func() Enumerator[Grouping[T, K]] {
		return &groupByEnumerator[T, K]{
			src:  Enumerable[T](src),
			ksel: keySelector,
		}
	}
}

func (e *groupByEnumerator[T, K]) Next() (def Grouping[T, K], _ error) {
	if e.ks == nil {
		ks := make([]K, 0)
		m := make(map[K][]T)
		err := ForEach(e.src, func(v T) error {
			k, err := e.ksel(v)
			if err != nil {
				return err
			}
			if _, ok := m[k]; !ok {
				ks = append(ks, k)
			}
			m[k] = append(m[k], v)
			return nil
		})
		if err != nil {
			return def, err
		}
		e.ks = ks
		e.m = m
	}

	if e.i >= len(e.ks) {
		return def, EOC
	}

	k := e.ks[e.i]
	e.i++
	return Grouping[T, K]{
		Enumerable: FromSlice(e.m[k]),
		Key:        k,
	}, nil
}
