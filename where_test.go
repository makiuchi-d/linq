package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestWhere(t *testing.T) {
	e1 := linq.FromSlice([]int16{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	e2 := linq.Where(e1, func(v int16) (bool, error) { return v%2 != 0, nil })
	r, err := linq.ToSlice(e2)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []int16{9, 7, 5, 3, 1}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("wants: %#v, got %#v", exp, r)
	}
}
