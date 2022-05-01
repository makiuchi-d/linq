package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestSum(t *testing.T) {
	src := []int16{100, 200, 300, 400, 500}
	r, err := linq.Sum(linq.FromSlice(src))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := 1500
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}

	r, err = linq.SumByFunc(linq.FromSlice(src), func(v int16) (int, error) {
		return int(v * 10), nil
	})
	exp *= 10
	if err != nil {
		t.Fatalf("%v", err)
	}
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}

	rf, err := linq.Sumf(
		linq.Select(linq.FromSlice(src), func(v int16) (float32, error) { return float32(v) / 16, nil }))
	if err != nil {
		t.Fatalf("%v", err)
	}
	expf := float64(1500) / 16
	if rf != expf {
		t.Fatalf("%v, wants %v", r, exp)
	}
	rf, err = linq.SumByFuncf(
		linq.Select(linq.FromSlice(src), func(v int16) (float32, error) { return float32(v), nil }),
		func(v float32) (float32, error) { return v / 16, nil })
	if err != nil {
		t.Fatalf("%v", err)
	}
	if rf != expf {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
