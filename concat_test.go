package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestConcat(t *testing.T) {
	e1 := linq.FromSlice([]int{4, 5, 6})
	e2 := linq.FromSlice([]int{1, 2, 3})

	r, err := linq.ToSlice(
		linq.Concat(e1, e2))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []int{4, 5, 6, 1, 2, 3}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
