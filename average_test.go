package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestAverage(t *testing.T) {
	src := linq.FromSlice([]byte{100, 110, 120, 130})
	r, err := linq.Average(src)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := float64(100+110+120+130) / 4
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
