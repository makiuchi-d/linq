package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestTakeWhile(t *testing.T) {
	src := linq.Range(0, 10)
	r, err := linq.ToSlice(
		linq.TakeWhile(src, func(n int) (bool, error) { return n < 5, nil }))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
