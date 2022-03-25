package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestSelectMany(t *testing.T) {
	s := [][]string{
		{"a", "bb", "ccc"},
		{"dddd", "eeeee", "ffffff"},
		{"ggggggg", "hhhhhhhh", "iiiiiiiii"},
	}

	e1 := linq.FromSlice(s)
	e2 := linq.SelectMany(e1,
		func(e []string) (linq.Enumerator[string], error) { return linq.FromSlice(e), nil },
		func(v string) (int, error) { return len(v), nil })

	r, err := linq.ToSlice(e2)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
