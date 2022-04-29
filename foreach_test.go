package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestForEach(t *testing.T) {
	c := 0
	s := 0
	err := linq.ForEach(
		linq.FromSlice([]int{1, 2, 3, 4, 5}),
		func(n int) error {
			c++
			s += n
			return nil
		})
	if err != nil {
		t.Fatalf("%v", err)
	}
	if c != 5 || s != 15 {
		t.Fatalf("(c, s) = (%v, %v) wants (%v, %v)", c, s, 5, 15)
	}
}
