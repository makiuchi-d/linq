package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
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

func TestExceptBy(t *testing.T) {
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

	e := linq.ExceptBy(
		linq.FromSlice(p1),
		linq.FromSlice(p2),
		func(p Planet) (string, error) { return p.Name, nil })

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []Planet{{"Venus", Rock, 2}}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
