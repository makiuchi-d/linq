package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestChunk(t *testing.T) {

	src := []int{1, 2, 3, 4, 5, 6, 7, 8}
	e := linq.Chunk(linq.FromSlice(src), 3)
	r, err := linq.ToSlice(linq.Select(e, linq.ToSlice[int]))
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}

	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
