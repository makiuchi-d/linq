package linq

import "sort"

type orderByFuncEnumerator[T any] struct {
	src    Enumerator[T]
	less   func(a, b T) bool
	sorted []T
	i      int
}

// OrderByFunc sorts the elements of a sequence by the provided less function.
func OrderByFunc[T any, E IEnumerable[T]](src E, less func(a, b T) bool) Enumerable[T] {
	return func() Enumerator[T] {
		return &orderByFuncEnumerator[T]{src: src(), less: less}
	}
}

func (o *orderByFuncEnumerator[T]) Next() (def T, _ error) {
	if o.sorted == nil {
		s, err := toSlice(o.src)
		if err != nil {
			return def, err
		}

		sort.Slice(s, func(i, j int) bool {
			return o.less(s[i], s[j])
		})

		o.sorted = s
	}

	if o.i >= len(o.sorted) {
		return def, EOC
	}

	i := o.i
	o.i++
	return o.sorted[i], nil
}
