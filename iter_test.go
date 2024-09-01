//go:build go1.23

package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestIterAll(t *testing.T) {
	src := linq.FromSlice([]int{1, 2, 3, 4, 5})
	c := 0
	s := 0
	for n, err := range src.All() {
		if err != nil {
			t.Fatalf("%v", err)
		}
		c++
		s += n
	}
	if c != 5 || s != 15 {
		t.Fatalf("(c, s) = (%v, %v) wants (%v, %v)", c, s, 5, 15)
	}
}
