package linq_test

import (
	"math"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestMax(t *testing.T) {
	src := []byte{
		189, 47, 155, 170, 155, 99, 136, 3, 161, 231,
	}

	r, err := linq.Max(linq.FromSlice(src))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := byte(231)
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}

	r, err = linq.MaxBy(linq.FromSlice(src), func(v byte) (float64, error) {
		return math.Abs(float64(v) - 150), nil
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp = byte(3)
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}

	r, err = linq.MaxByFunc(linq.FromSlice(src), func(a, b byte) (bool, error) {
		if a%2 != b%2 {
			return a%2 == 0, nil
		}
		return a > b, nil
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp = byte(170)
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
