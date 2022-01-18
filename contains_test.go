package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestContains(t *testing.T) {
	r, err := linq.Contains(linq.FromSlice([]int{1, 2, 3, 4, 5}), 3)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if r != true {
		t.Fatalf("3 must be contained")
	}
	r, err = linq.Contains(linq.FromSlice([]int{1, 2, 3, 4, 5}), 6)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if r != false {
		t.Fatalf("6 must not be contained")
	}
}

func TestContainsFunc(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}
	src := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
		{Name: "lemon", Code: 12},
	}
	e := linq.FromSlice(src)
	v := Product{Name: "apple", Code: 1}
	r, err := linq.ContainsFunc(e, v, func(a, b Product) (bool, error) {
		return a.Name == b.Name, nil
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	if r != true {
		t.Fatalf("must be true")
	}
}
