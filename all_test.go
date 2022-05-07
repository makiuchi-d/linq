package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestAll(t *testing.T) {
	isEven := func(v int) (bool, error) {
		return v%2 == 0, nil
	}

	tests := []struct {
		src []int
		exp bool
	}{
		{[]int{2, 4, 6, 8, 10}, true},
		{[]int{2, 4, 6, 7, 10}, false},
	}

	for _, test := range tests {
		e := linq.FromSlice(test.src)
		r, err := linq.All(e, isEven)
		if err != nil {
			t.Fatalf("%v", err)
		}
		if r != test.exp {
			t.Fatalf("wants: %v, got %v", test.exp, r)
		}
	}
}
