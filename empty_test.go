package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestEmpty(t *testing.T) {
	r, err := linq.ToSlice(linq.Empty[uint]())
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(r) != 0 {
		t.Fatalf("not empty: %#v", r)
	}
	if reflect.TypeOf(r) != reflect.TypeOf([]uint{}) {
		t.Fatalf("%T != []uint", r)
	}
}
