package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestAny(t *testing.T) {
	isOdd := func(v int) (bool, error) {
		return v%2 != 0, nil
	}

	tests := []struct {
		src []int
		exp bool
	}{
		{[]int{2, 4, 6, 8, 10}, false},
		{[]int{2, 4, 6, 7, 10}, true},
	}

	for _, test := range tests {
		e := linq.FromSlice(test.src)
		r, err := linq.Any(e, isOdd)
		if err != nil {
			t.Fatalf("%v", err)
		}
		if r != test.exp {
			t.Fatalf("wants: %v, got %v", test.exp, r)
		}
	}
}
