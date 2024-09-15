//go:build go1.23

package linq

import "iter"

// All returns an iterator function
func (e Enumerable[T]) All() iter.Seq2[T, error] {
	er := e()
	return func(yield func(T, error) bool) {
		for {
			v, err := er.Next()
			if err != nil {
				if isEOC(err) {
					return
				}
				yield(v, err)
				return
			}

			if !yield(v, nil) {
				return
			}
		}
	}
}

// All returns an iterator function
func (e OrderedEnumerable[T]) All() iter.Seq2[T, error] {
	er := e()
	return func(yield func(T, error) bool) {
		for {
			v, err := er.Next()
			if err != nil {
				if isEOC(err) {
					return
				}
				yield(v, err)
				return
			}

			if !yield(v, nil) {
				return
			}
		}
	}
}
