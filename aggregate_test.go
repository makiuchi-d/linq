package linq_test

import (
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestAggregate(t *testing.T) {
	src := []byte{100, 100, 200, 200}

	r, err := linq.Aggregate(
		linq.FromSlice(src),
		0,
		func(acc int, v byte) (int, error) {
			return acc + int(v), nil
		})
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := 600
	if r != exp {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
