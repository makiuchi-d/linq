package linq

type joinEnumerator[T1, T2, U any, K comparable] struct {
	eOut  Enumerator[T1]
	eIn   Enumerator[T2]
	ksOut func(T1) (K, error)
	ksIn  func(T2) (K, error)
	rSel  func(T1, T2) (U, error)

	t1  *T1
	mt2 map[K][]T2
	i   int
}

// Join correlates the elements of two sequences based on matching keys.
func Join[T1, T2, U any, K comparable](
	outer Enumerator[T1],
	inner Enumerator[T2],
	outerKeySelector func(T1) (K, error),
	innerKeySelector func(T2) (K, error),
	resultSelector func(T1, T2) (U, error),
) Enumerator[U] {
	return &joinEnumerator[T1, T2, U, K]{
		eOut:  outer,
		eIn:   inner,
		ksOut: outerKeySelector,
		ksIn:  innerKeySelector,
		rSel:  resultSelector,
	}
}

func (e *joinEnumerator[T1, T2, U, K]) Next() (def U, _ error) {
	if e.t1 == nil {
		t1, err := e.eOut.Next()
		if err != nil {
			return def, err
		}
		e.t1 = &t1
	}

	if e.mt2 == nil {
		m, err := innerMap(e.eIn, e.ksIn)
		if err != nil {
			return def, err
		}
		e.mt2 = m
	}

	k, err := e.ksOut(*e.t1)
	if err != nil {
		return def, err
	}

	s := e.mt2[k]

	if e.i >= len(s) {
		e.i = 0
		e.t1 = nil
		return e.Next()
	}

	i := e.i
	e.i++

	return e.rSel(*e.t1, s[i])
}

func innerMap[T2 any, K comparable](e Enumerator[T2], ks func(T2) (K, error)) (map[K][]T2, error) {
	m := make(map[K][]T2)
	for {
		t2, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return m, nil
			}
			return nil, err
		}

		k, err := ks(t2)
		if err != nil {
			return nil, err
		}
		m[k] = append(m[k], t2)
	}
}
