package linq_test

import (
	"errors"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestLast(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7}
	r, err := linq.Last(linq.FromSlice(src),
		func(v int) (bool, error) {
			return v%2 == 0, nil
		})
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := 6
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}

	_, err = linq.Last(linq.FromSlice(src),
		func(v int) (bool, error) { return false, nil })
	if !errors.Is(err, linq.InvalidOperation) {
		t.Fatalf("%#v, wants %#v", err, linq.InvalidOperation)
	}
}

func TestLastOrDefault(t *testing.T) {
	src := []int{1, 2, 3}
	def := 42
	r, err := linq.LastOrDefault(linq.FromSlice(src),
		func(v int) (bool, error) { return false, nil },
		def)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := def
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
	r, err = linq.LastOrDefault(linq.FromSlice(src),
		func(v int) (bool, error) { return true, nil },
		def)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp = 3
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
