package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestIntersect(t *testing.T) {
	fst := linq.FromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	snd := linq.FromSlice([]int{0, 2, 4, 6, 8, 10})

	e := linq.Intersect(
		fst, snd,
		func(a, b int) (bool, error) { return a == b, nil },
		func(a int) (int, error) { return a / 3, nil })

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []int{2, 4, 6}

	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
