package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestOrderBy(t *testing.T) {
	fruits := []string{
		"grape", "passionfruit", "banana", "mango",
		"orange", "raspberry", "apple", "blueberry"}

	e1 := linq.FromSlice(fruits)
	e2 := linq.OrderBy(e1, func(v string) (int, error) { return len(v), nil })
	r, err := linq.ToSlice(e2)
	if err != nil {
		t.Fatalf("%v", err)
	}

	for i := 1; i < len(r); i++ {
		if len(r[i-1]) > len(r[i]) {
			t.Fatalf("len(%q) > len(%q)\n\t%v", r[i-1], r[i], r)
		}
	}
}

func TestOrderByDescending(t *testing.T) {
	fruits := []string{
		"grape", "passionfruit", "banana", "mango",
		"orange", "raspberry", "apple", "blueberry"}

	e1 := linq.FromSlice(fruits)
	e2 := linq.OrderByDescending(e1, func(v string) (int, error) { return len(v), nil })
	r, err := linq.ToSlice(e2)
	if err != nil {
		t.Fatalf("%v", err)
	}

	for i := 1; i < len(r); i++ {
		if len(r[i-1]) < len(r[i]) {
			t.Fatalf("len(%q) < len(%q)\n\t%v", r[i-1], r[i], r)
		}
	}
}

func TestThenBy(t *testing.T) {
	fruits := []string{
		"grape", "passionfruit", "banana", "mango",
		"orange", "raspberry", "apple", "blueberry"}

	e1 := linq.FromSlice(fruits)
	e2 := linq.OrderByDescending(e1, func(v string) (int, error) { return len(v), nil })
	e3 := linq.ThenBy(e2, func(v string) (string, error) { return v, nil })
	r, err := linq.ToSlice(e3)
	if err != nil {
		t.Fatalf("%v", err)
	}

	for i := 1; i < len(r); i++ {
		if len(r[i-1]) < len(r[i]) {
			t.Fatalf("len(%q) < len(%q)\n\t%v", r[i-1], r[i], r)
		}
		if len(r[i-1]) == len(r[i]) && r[i-1] > r[i] {
			t.Fatalf("%q > %q\n\t%v", r[i-1], r[i], r)
		}
	}
}

func TestThenByDescending(t *testing.T) {
	fruits := []string{
		"grape", "passionfruit", "banana", "mango",
		"orange", "raspberry", "apple", "blueberry"}

	e1 := linq.FromSlice(fruits)
	e2 := linq.OrderBy(e1, func(v string) (int, error) { return len(v), nil })
	e3 := linq.ThenByDescending(e2, func(v string) (string, error) { return v, nil })
	r, err := linq.ToSlice(e3)
	if err != nil {
		t.Fatalf("%v", err)
	}

	for i := 1; i < len(r); i++ {
		if len(r[i-1]) > len(r[i]) {
			t.Fatalf("len(%q) > len(%q)\n\t%v", r[i-1], r[i], r)
		}
		if len(r[i-1]) == len(r[i]) && r[i-1] < r[i] {
			t.Fatalf("%q < %q\n\t%v", r[i-1], r[i], r)
		}
	}
}
