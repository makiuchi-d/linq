//go:build go1.23

package linq_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestIterAll(t *testing.T) {
	src := linq.FromSlice([]int{1, 2, 3, 4, 5})
	c := 0
	s := 0
	for n, err := range src.All() {
		if err != nil {
			t.Fatalf("%v", err)
		}
		c++
		s += n
	}
	if c != 5 || s != 15 {
		t.Fatalf("(c, s) = (%v, %v) wants (%v, %v)", c, s, 5, 15)
	}
}

func TestFromIterator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	e := linq.FromIterator(ctx, func(yield func(int) bool) {
		_ = yield(1) && yield(2) && yield(3) && yield(4) && yield(5)
	})
	e = linq.Take(e, 3)

	a, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	exp := []int{1, 2, 3}
	if !reflect.DeepEqual(a, exp) {
		t.Fatalf("ToSlice: %v, wants %v", a, exp)
	}
}

func TestFromIterator2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	e := linq.FromIterator2(ctx, func(yield func(int, string) bool) {
		_ = yield(1, "one") && yield(2, "two") && yield(3, "three")
	})

	a, err := linq.ToMap(e)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	exp := map[int]string{1: "one", 2: "two", 3: "three"}
	if !reflect.DeepEqual(a, exp) {
		t.Fatalf("ToMap: %v, wants %v", a, exp)
	}
}
