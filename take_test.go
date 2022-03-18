package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestTake(t *testing.T) {
	tests := []struct {
		src linq.Enumerator[int]
		cnt int
		exp []int
	}{
		{linq.Range(0, 5), 3, []int{0, 1, 2}},
		{linq.Range(0, 5), 8, []int{0, 1, 2, 3, 4}},
	}
	for _, test := range tests {
		r, err := linq.ToSlice(
			linq.Take(test.src, test.cnt))
		if err != nil {
			t.Fatalf("%v: %v", test.cnt, err)
		}
		if !reflect.DeepEqual(r, test.exp) {
			t.Fatalf("%v, wants %v", r, test.exp)
		}
	}

}
