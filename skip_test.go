package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestSkip(t *testing.T) {
	tests := []struct {
		src linq.Enumerator[int]
		n   int
		exp []int
	}{
		{linq.Range(0, 5), 0, []int{0, 1, 2, 3, 4}},
		{linq.Range(0, 5), 2, []int{2, 3, 4}},
		{linq.Range(0, 5), 5, []int{}},
		{linq.Range(0, 5), 6, []int{}},
	}
	for _, test := range tests {
		e := linq.Skip(test.src, test.n)
		r, err := linq.ToSlice(e)
		if err != nil {
			t.Fatalf("Skip(%v): %v", test.n, err)
		}
		if !reflect.DeepEqual(r, test.exp) {
			t.Fatalf("Skip(%v): %v, wants %v", test.n, r, test.exp)
		}
	}
}
