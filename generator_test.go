package linq_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func ExampleGenerator() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gen := linq.Generator(ctx, func(yield func(int) error) error {
		for i := 0; i < 10; i++ {
			err := yield(i)
			if err != nil {
				return err
			}
		}
		return nil
	})

	r, err := linq.ToSlice(gen)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func TestGenerator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gen := linq.Generator(ctx, func(yield func(int) error) error {
		for i := 0; i < 10; i++ {
			err := yield(i)
			if err != nil {
				return err
			}
		}
		return nil
	})

	r, err := linq.ToSlice(gen)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("wants: %#v, got %#v", r, exp)
	}
}

func TestGeneratorError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	experr := errors.New("expect error")

	gen := linq.Generator(ctx, func(yield func(int) error) error {
		err := yield(1)
		if err != nil {
			return err
		}
		return experr
	})

	g := gen()
	n, err := g.Next()
	if n != 1 || err != nil {
		t.Fatalf("Next: %v, %v, wants 1, nil", n, err)
	}
	_, err = g.Next()
	if !errors.Is(err, experr) {
		t.Fatalf("error: %q, wants %q", err, experr)
	}

	_, err = g.Next()
	if !errors.Is(err, experr) {
		t.Fatalf("error: %q, wants %q", err, experr)
	}
}

func TestGeneratorCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	gen := linq.Generator(ctx, func(yield func(int) error) error {
		for i := 1; i <= 3; i++ {
			err := yield(i)
			if err != nil {
				return err
			}
		}
		return nil
	})

	g := gen()
	n, err := g.Next()
	if n != 1 || err != nil {
		t.Fatalf("Next: %v, %v, wants 1, nil", n, err)
	}
	cancel()
	_, err = g.Next()
	if err == nil || !errors.Is(err, ctx.Err()) {
		t.Fatalf("error: %q, wants \"context canceled\"", err)
	}

	_, err = g.Next()
	if err == nil || !errors.Is(err, ctx.Err()) {
		t.Fatalf("error: %q, wants \"context canceled\"", err)
	}
}

func TestGeneratorPanic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ep := errors.New("expect panic")

	gen := linq.Generator(ctx, func(yield func(int) error) error {
		err := yield(1)
		if err != nil {
			return err
		}
		panic(ep)
	})

	g := gen()
	n, err := g.Next()
	if n != 1 || err != nil {
		t.Fatalf("Next: %v, %v, wants 1, nil", n, err)
	}

	defer func() {
		p := recover()
		if p != ep {
			t.Fatalf("panic: %v, wants %v", p, ep)
		}
	}()

	_, _ = g.Next()

	t.Fatalf("no panic")
}
