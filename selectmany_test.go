package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestSelectMany(t *testing.T) {
	s := [][]string{
		{"a", "bb", "ccc"},
		{"dddd", "eeeee", "ffffff"},
		{"ggggggg", "hhhhhhhh", "iiiiiiiii"},
	}

	e1 := linq.FromSlice(s)
	e2 := linq.SelectMany(e1,
		func(e []string) (linq.OrderedEnumerable[string], error) {
			return linq.OrderBy(linq.FromSlice(e), func(s string) (string, error) { return s, nil }), nil
		},
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
