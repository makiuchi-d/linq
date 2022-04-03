package linq_test

import (
	"errors"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestFirst(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	r, err := linq.First(linq.FromSlice(src),
		func(v int) (bool, error) {
			return (v > 5 && v%3 == 0), nil
		})
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := 6
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}

	_, err = linq.First(linq.FromSlice(src),
		func(v int) (bool, error) { return false, nil })
	if !errors.Is(err, linq.InvalidOperation) {
		t.Fatalf("%#v, wants %#v", err, linq.InvalidOperation)
	}

}

func TestFirstOrDefault(t *testing.T) {
	src := []int{1, 2, 3}
	def := 42
	r, err := linq.FirstOrDefault(linq.FromSlice(src),
		func(v int) (bool, error) { return false, nil },
		def)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := def
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}

	r, err = linq.FirstOrDefault(linq.FromSlice(src),
		func(v int) (bool, error) { return true, nil },
		def)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp = 1
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
