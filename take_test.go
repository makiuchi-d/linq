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

func TestTakeWhile(t *testing.T) {
	src := linq.Range(0, 10)
	r, err := linq.ToSlice(
		linq.TakeWhile(src, func(n int) (bool, error) { return n < 5, nil }))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}

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
