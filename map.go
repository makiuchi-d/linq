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
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return &mapEnumerator[K, V]{m: m, k: keys}
}

// ToMap creates a map[K]V from an Enumerator[T].
// T must be a type KeyValue[K, V].
func ToMap[K comparable, V any](src Enumerator[KeyValue[K, V]]) (map[K]V, error) {
	m := make(map[K]V)
	for {
		kv, err := src.Next()
		if err != nil {
			if isEOC(err) {
				break
			}
			return m, err
		}

		m[kv.Key] = kv.Value
	}
	return m, nil
}

func (e *mapEnumerator[K, V]) Next() (KeyValue[K, V], error) {
	if e.i >= len(e.k) {
		var kv KeyValue[K, V]
		return kv, EOC
	}
	k := e.k[e.i]
	e.i++
	return KeyValue[K, V]{Key: k, Value: e.m[k]}, nil
}
