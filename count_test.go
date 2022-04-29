package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestCount(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	r, err := linq.Count(linq.FromSlice(src))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := 9
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
