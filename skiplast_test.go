package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestSkipLast(t *testing.T) {
	src := linq.Range(0, 10)
	r, err := linq.ToSlice(
		linq.SkipLast(src, 3))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []int{0, 1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
