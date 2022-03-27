package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestReverse(t *testing.T) {
	src := linq.FromSlice([]int{1, 2, 3, 4, 5})
	r, err := linq.ToSlice(
		linq.Reverse(src))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
