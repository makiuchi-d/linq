package linq

import "errors"

func isEOC(err error) bool {
	return errors.Is(err, EOC)
}

func isOutOfRange(err error) bool {
	return errors.Is(err, OutOfRange)
}

func isInvalidOperation(err error) bool {
	return errors.Is(err, InvalidOperation)
}

type hashMap[H comparable, V any] struct {
	m    map[H][]V
	hash func(V) (H, error)
	eq   func(V, V) (bool, error)
}

func newHashMap[H comparable, V any](hash func(V) (H, error), eq func(V, V) (bool, error)) *hashMap[H, V] {
	return &hashMap[H, V]{
		m:    make(map[H][]V),
		hash: hash,
		eq:   eq,
	}
}

func newKeyMap[K comparable, V any](ks func(V) (K, error)) *hashMap[K, V] {
	return newHashMap(ks, alwaysEqual[V])
}

func alwaysEqual[T any](_, _ T) (bool, error) {
	return true, nil
}

func alwaysNotEqual[T any](_, _ T) (bool, error) {
	return false, nil
}

func (hm *hashMap[H, V]) gets(h H) []V {
	return hm.m[h]
}

func (hm *hashMap[H, V]) has(v V) (bool, error) {
	h, err := hm.hash(v)
	if err != nil {
		return false, err
	}
	for _, u := range hm.m[h] {
		eq, err := hm.eq(u, v)
		if err != nil {
			return false, err
		}
		if eq {
			return true, nil
		}
	}
	return false, nil
}

func (hm *hashMap[H, V]) add(v V) (bool, error) {
	h, err := hm.hash(v)
	if err != nil {
		return false, err
	}
	for _, u := range hm.m[h] {
		eq, err := hm.eq(u, v)
		if err != nil {
			return false, err
		}
		if eq {
			return false, nil
		}
	}
	hm.m[h] = append(hm.m[h], v)
	return true, nil
}

func (hm *hashMap[H, V]) addAll(e Enumerator[V]) error {
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return nil
			}
			return err
		}

		if _, err = hm.add(v); err != nil {
			return err
		}
	}
}
