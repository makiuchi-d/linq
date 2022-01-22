package linq_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestZip(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	words := []string{"one", "two", "three"}

	e := linq.Zip(
		linq.FromSlice(nums),
		linq.FromSlice(words),
		func(n int, w string) (string, error) {
			return fmt.Sprintf("%v %v", n, w), nil
		})

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []string{"1 one", "2 two", "3 three"}
	if !reflect.DeepEqual(exp, r) {
		t.Fatalf("%q, wants %q", r, exp)
	}
}
