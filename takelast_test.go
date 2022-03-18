package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestTakeLast(t *testing.T) {
	tests := []struct {
		src linq.Enumerator[int]
		n   int
		exp []int
	}{
		{linq.Range(0, 10), 5, []int{5, 6, 7, 8, 9}},
		{linq.Range(0, 3), 5, []int{0, 1, 2}},
	}
	for i, test := range tests {
		r, err := linq.ToSlice(
			linq.TakeLast(test.src, test.n))
		if err != nil {
			t.Fatalf("%v: %v", i, err)
		}
		if !reflect.DeepEqual(r, test.exp) {
			t.Fatalf("%v, wants %v", r, test.exp)
		}
	}
}
