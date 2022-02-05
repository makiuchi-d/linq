package linq

import "errors"

func isEOC(err error) bool {
	return errors.Is(err, EOC)
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

func (hm *hashMap[K, V]) add(v V) (bool, error) {
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
