package linq

import (
	"context"
)

type generator[T any] struct {
	ch     <-chan T
	resume func() error
	err    error
	panic  any
}

var _ Enumerator[int] = &generator[int]{}

// Generator makes the given function an enumerator generator.
func Generator[T any](ctx context.Context, g func(yield func(T) error) error) Enumerable[T] {
	return func() Enumerator[T] {
		do := make(chan struct{})
		done := make(chan struct{})
		ch := make(chan T)

		ge := &generator[T]{ch: ch}

		ge.resume = func() error {
			select {
			case <-done:
				return ge.err
			case do <- struct{}{}:
				return nil
			}
		}

		go func() {
			defer close(done)
			defer close(ch)
			defer func() {
				ge.panic = recover()
			}()

			select {
			case <-ctx.Done():
				ge.err = ctx.Err()
				return
			case <-do:
			}

			yield := func(t T) error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case ch <- t:
				}

				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-do:
					return nil
				}
			}

			ge.err = g(yield)
			if ge.err == nil {
				ge.err = EOC
			}
		}()

		return ge
	}
}

func (g *generator[T]) Next() (zero T, _ error) {
	if err := g.resume(); err != nil {
		return zero, err
	}
	t, ok := <-g.ch
	if !ok {
		if g.panic != nil {
			panic(g.panic)
		}
		return zero, g.err
	}
	return t, nil
}
