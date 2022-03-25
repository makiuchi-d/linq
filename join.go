package linq

type joinEnumerator[S1, S2, T any, K comparable] struct {
	eOut  Enumerator[S1]
	eIn   Enumerator[S2]
	ksOut func(S1) (K, error)
	ksIn  func(S2) (K, error)
	rSel  func(S1, S2) (T, error)

	s1  *S1
	ks1 K
	ms2 map[K][]S2
	i   int
}

// Join correlates the elements of two sequences based on matching keys.
func Join[S1, S2, T any, K comparable](
	outer Enumerator[S1],
	inner Enumerator[S2],
	outerKeySelector func(S1) (K, error),
	innerKeySelector func(S2) (K, error),
	resultSelector func(S1, S2) (T, error),
) Enumerator[T] {
	return &joinEnumerator[S1, S2, T, K]{
		eOut:  outer,
		eIn:   inner,
		ksOut: outerKeySelector,
		ksIn:  innerKeySelector,
		rSel:  resultSelector,
	}
}

func (e *joinEnumerator[S1, S2, T, K]) Next() (def T, _ error) {
	if e.s1 == nil {
		s1, err := e.eOut.Next()
		if err != nil {
			return def, err
		}
		ks1, err := e.ksOut(s1)
		if err != nil {
			return def, err
		}
		e.s1 = &s1
		e.ks1 = ks1
	}

	if e.ms2 == nil {
		m, err := innerMap(e.eIn, e.ksIn)
		if err != nil {
			return def, err
		}
		e.ms2 = m
	}

	s := e.ms2[e.ks1]

	if e.i >= len(s) {
		e.i = 0
		e.s1 = nil
		return e.Next()
	}

	i := e.i
	e.i++

	return e.rSel(*e.s1, s[i])
}

func innerMap[S2 any, K comparable](e Enumerator[S2], ks func(S2) (K, error)) (map[K][]S2, error) {
	m := make(map[K][]S2)
	for {
		s2, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return m, nil
			}
			return nil, err
		}

		k, err := ks(s2)
		if err != nil {
			return nil, err
		}
		m[k] = append(m[k], s2)
	}
}
