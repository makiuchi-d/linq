package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestDistinct(t *testing.T) {
	ages := []int{21, 46, 46, 55, 17, 21, 55, 55, 25, 25}

	e := linq.Distinct(
		linq.FromSlice(ages),
		func(a, b int) (bool, error) { return a == b, nil },
		func(a int) (int, error) { return a / 10, nil })

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []int{21, 46, 55, 17, 25}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}

func TestDistinctBy(t *testing.T) {
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
	planets := []Planet{
		{"Marcury", Rock, 1},
		{"Venus", Rock, 2},
		{"Earth", Rock, 3},
		{"Mars", Rock, 4},
		{"Jupiter", Gas, 5},
		{"Saturn", Gas, 6},
		{"Uranus", Liquid, 7},
		{"Neptune", Liquid, 8},
		{"Pluto", Ice, 9}, // dwarf planet
	}

	e := linq.DistinctBy(linq.FromSlice(planets), func(p Planet) (PlanetType, error) {
		return p.Type, nil
	})
	r, err := linq.ToSlice(linq.Select(e, func(p Planet) (string, error) { return p.Name, nil }))
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []string{"Marcury", "Jupiter", "Uranus", "Pluto"}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
