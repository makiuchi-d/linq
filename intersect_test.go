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

func TestIntersectBy(t *testing.T) {
	type PlanetType int
	const (
		Rock PlanetType = iota
		Ice
		Gas
		Liquid
	)
	type Planet struct {
		Name         string
		Type         PlanetType
		OrderFromSun int
	}
	p1 := []Planet{
		{"Marcury", Rock, 1},
		{"Venus", Rock, 2},
		{"Earth", Rock, 3},
		{"Jupiter", Gas, 5},
	}
	p2 := []Planet{
		{"Marcury", Rock, 1},
		{"Earth", Rock, 3},
		{"Mars", Rock, 4},
		{"Jupiter", Gas, 5},
	}

	e := linq.IntersectBy(
		linq.FromSlice(p1),
		linq.FromSlice(p2),
		func(p Planet) (string, error) { return p.Name, nil })

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []Planet{
		{"Marcury", Rock, 1},
		{"Earth", Rock, 3},
		{"Jupiter", Gas, 5},
	}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
