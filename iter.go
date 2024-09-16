//go:build go1.23

package linq

import (
	"context"
	"iter"
)

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

// FromIterator generates an IEnumerable[T] from an iterator function.
func FromIterator[T any](ctx context.Context, it iter.Seq[T]) Enumerable[T] {
	return Generator(ctx, func(yield func(T) error) error {
		for v := range it {
			if err := yield(v); err != nil {
				return err
			}
		}
		return nil
	})
}

// FromIterator2 generates an IEnumerable[KeyValue[K,V]] from an iterator function.
func FromIterator2[K comparable, V any](ctx context.Context, it iter.Seq2[K, V]) Enumerable[KeyValue[K, V]] {
	return Generator(ctx, func(yield func(KeyValue[K, V]) error) error {
		for k, v := range it {
			if err := yield(KeyValue[K, V]{k, v}); err != nil {
				return err
			}
		}
		return nil
	})
}
