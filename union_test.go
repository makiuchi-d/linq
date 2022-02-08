package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestUnion(t *testing.T) {
	fst := linq.FromSlice([]int{2, 3, 4, 5})
	snd := linq.FromSlice([]int{0, 2, 4, 6, 8})

	e := linq.Union(
		fst, snd,
		func(a, b int) (bool, error) { return a == b, nil },
		func(a int) (int, error) { return a / 3, nil })

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []int{2, 3, 4, 5, 0, 6, 8}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
