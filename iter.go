//go:build go1.23

package linq

import "iter"

func (e Enumerable[T]) All() iter.Seq2[T, error] {
	return iterAll(e())
}

func (e OrderedEnumerable[T]) All() iter.Seq2[T, error] {
	return iterAll(e())
}

func iterAll[T any, E Enumerator[T]](e E) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for {
			v, err := e.Next()
			if err != nil {
				if isEOC(err) {
					return
				}
			}

			if !yield(v, err) {
				return
			}

			if err != nil {
				return
			}
		}
	}
}
