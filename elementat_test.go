package linq_test

import (
	"errors"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestElementAt(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}

	r, err := linq.ElementAt(linq.FromSlice(src), 3)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := 4
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
	_, err = linq.ElementAt(linq.FromSlice(src), 10)
	if !errors.Is(err, linq.OutOfRange) {
		t.Fatalf("%#v, wants %#v", err, linq.OutOfRange)
	}

	r, err = linq.ElementAtOrDefault(linq.FromSlice(src), 10)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if r != 0 {
		t.Fatalf("%v, wants 0", r)
	}
}
