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

type keyMap[K comparable, V any] struct {
	m  map[K]struct{}
	ks func(v V) (K, error)
}

func newKeyMap[K comparable, V any](ks func(V) (K, error)) *keyMap[K, V] {
	return &keyMap[K, V]{
		m:  make(map[K]struct{}),
		ks: ks,
	}
}

func (km *keyMap[K, V]) has(v V) (bool, error) {
	k, err := km.ks(v)
	if err != nil {
		return false, err
	}
	_, ok := km.m[k]
	return ok, nil
}

func (km *keyMap[K, V]) add(v V) (bool, error) {
	k, err := km.ks(v)
	if err != nil {
		return false, err
	}
	_, ok := km.m[k]
	if !ok {
		km.m[k] = struct{}{}
	}
	return !ok, nil
}

func (km *keyMap[K, V]) addAll(e Enumerator[V]) error {
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return nil
			}
			return err
		}
		k, err := km.ks(v)
		if err != nil {
			return err
		}
		km.m[k] = struct{}{}
	}
}
