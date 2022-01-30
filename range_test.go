package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestRange(t *testing.T) {
	e := linq.Range(3, 5)
	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []int{3, 4, 5, 6, 7}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
