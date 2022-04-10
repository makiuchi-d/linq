package linq

// KeyValue pair as an element of map[K]V
type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

type mapEnumerator[K comparable, V any] struct {
	m map[K]V
	k []K
	i int
}

// FromMap generates an Enumerator[T] from a map.
func FromMap[T ~map[K]V, K comparable, V any](m T) Enumerator[KeyValue[K, V]] {
	return &mapEnumerator[K, V]{m: m}
}

func (e *mapEnumerator[K, V]) Next() (def KeyValue[K, V], _ error) {
	if e.k == nil {
		ks := make([]K, 0, len(e.m))
		for k := range e.m {
			ks = append(ks, k)
		}
		e.k = ks
	}
	if e.i >= len(e.k) {
		return def, EOC
	}
	k := e.k[e.i]
	e.i++
	return KeyValue[K, V]{Key: k, Value: e.m[k]}, nil
}

// ToMap creates a map[K]V from an Enumerator[T].
// T must be a type KeyValue[K, V].
func ToMap[K comparable, V any](src Enumerator[KeyValue[K, V]]) (map[K]V, error) {
	m := make(map[K]V)
	for {
		kv, err := src.Next()
		if err != nil {
			if isEOC(err) {
				return m, nil
			}
			return m, err
		}

		m[kv.Key] = kv.Value
	}
}

// ToMapFunc creates a map[K]V from an Enumerator[T] according to specified key-value selector function.
func ToMapFunc[T any, K comparable, V any](src Enumerator[T], selector func(T) (K, V, error)) (map[K]V, error) {
	m := make(map[K]V)
	for {
		t, err := src.Next()
		if err != nil {
			if isEOC(err) {
				return m, nil
			}
			return m, err
		}

		k, v, err := selector(t)
		if err != nil {
			return m, err
		}

		m[k] = v
	}
}
