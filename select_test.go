package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestSelect(t *testing.T) {
	e1 := linq.FromSlice([]string{
		"a", "bb", "ccc", "dddd", "eeeee",
	})
	e2 := linq.Select(e1, func(s string) (int, error) { return len(s), nil })

	r, err := linq.ToSlice(e2)
	if err != nil {
		t.Fatalf(err.Error())
	}

	exp := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("got %#v, wants %#v", r, exp)
	}
}
