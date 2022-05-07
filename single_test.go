package linq_test

import (
	"errors"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestSingle(t *testing.T) {
	src := linq.FromSlice([]int{1, 2, 3})

	r, err := linq.Single(linq.Where(src,
		func(v int) (bool, error) {
			return v%2 == 0, nil
		}))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := 2
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}

	_, err = linq.Single(linq.Where(src,
		func(v int) (bool, error) {
			return v%2 != 0, nil
		}))
	if !errors.Is(err, linq.InvalidOperation) {
		t.Fatalf("%#v, wants %#v", err, linq.InvalidOperation)
	}
}

func TestSingleOrDefault(t *testing.T) {
	src := linq.FromSlice([]int{1, 2, 3})
	def := 42
	r, err := linq.SingleOrDefault(linq.Where(src,
		func(v int) (bool, error) { return v%2 == 0, nil }),
		def)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := 2
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
	r, err = linq.SingleOrDefault(linq.Empty[int](), def)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp = def
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
	_, err = linq.SingleOrDefault(linq.Where(src,
		func(v int) (bool, error) { return v%2 != 0, nil }),
		def)
	if !errors.Is(err, linq.InvalidOperation) {
		t.Fatalf("%#v, wants %#v", err, linq.InvalidOperation)
	}
}
