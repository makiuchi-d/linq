package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestSkipWhile(t *testing.T) {
	src := linq.Range(0, 10)
	r, err := linq.ToSlice(
		linq.SkipWhile(src, func(n int) (bool, error) { return n < 5, nil }))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []int{5, 6, 7, 8, 9}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
