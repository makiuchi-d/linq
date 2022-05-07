package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestRepeat(t *testing.T) {
	r, err := linq.ToSlice(linq.Repeat(rune('A'), 5))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []rune{'A', 'A', 'A', 'A', 'A'}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
