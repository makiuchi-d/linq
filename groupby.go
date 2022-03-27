package linq

// Grouping represents a collection of objects that have a common key.
type Grouping[K comparable, T any] interface {
	Enumerator[T]
	Key() K
}

type grouping[K comparable, T any] struct {
	Enumerator[T]
	key K
}

func (g grouping[K, T]) Key() K {
	return g.key
}

type groupByEnumerator[T any, K comparable] struct {
	src  Enumerator[T]
	ksel func(T) (K, error)

	ks []K
	m  map[K][]T
	i  int
}

// GroupBy groups the elements of a sequence according to a specified key selector function.
func GroupBy[T any, K comparable](src Enumerator[T], keySelector func(T) (K, error)) Enumerator[Grouping[K, T]] {
	return &groupByEnumerator[T, K]{
		src:  src,
		ksel: keySelector,
	}
}

func (e *groupByEnumerator[T, K]) Next() (def Grouping[K, T], _ error) {
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
	return grouping[K, T]{
		Enumerator: FromSlice(e.m[k]),
		key:        k,
	}, nil
}
