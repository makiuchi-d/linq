package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestExcept(t *testing.T) {
	e1 := linq.FromSlice([]string{"Mercury", "Venus", "Earth", "Jupiter"})
	e2 := linq.FromSlice([]string{"Mercury", "Earth", "Mars", "Jupiter"})

	e := linq.Except(e1, e2,
		func(a, b string) (bool, error) { return a == b, nil },
		func(a string) (int, error) { return len(a), nil })

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []string{"Venus"}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%q, wants %q", r, exp)
	}
}
