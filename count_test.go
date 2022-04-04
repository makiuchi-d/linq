package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestCount(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	r, err := linq.Count(linq.FromSlice(src),
		func(v int) (bool, error) { return v%2 == 0, nil })
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := 4
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}

	r, err = linq.Count(linq.FromSlice(src),
		func(v int) (bool, error) { return true, nil })
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp = 9
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
