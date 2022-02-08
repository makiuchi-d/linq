package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestUnionBy(t *testing.T) {
	type Planet struct {
		Name         string
		OrderFromSun int
	}
	fst := []Planet{
		{"Marcury", 1},
		{"Venus", 2},
		{"Earth", 3},
		{"Mars", 4},
		{"Jupiter", 5},
	}
	snd := []Planet{
		{"Mars", 4},
		{"Jupiter", 5},
		{"Saturn", 6},
		{"Uranus", 7},
		{"Neptune", 8},
	}

	e := linq.UnionBy(
		linq.FromSlice(fst), linq.FromSlice(snd),
		func(v Planet) (int, error) { return v.OrderFromSun, nil })

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []Planet{
		{"Marcury", 1},
		{"Venus", 2},
		{"Earth", 3},
		{"Mars", 4},
		{"Jupiter", 5},
		{"Saturn", 6},
		{"Uranus", 7},
		{"Neptune", 8},
	}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
