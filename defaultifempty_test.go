package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestDefaultIfEmpty(t *testing.T) {
	src := linq.FromSlice([]string{"aaa", "bbb", "ccc"})
	r, err := linq.ToSlice(
		linq.DefaultIfEmpty(src, "default"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []string{"aaa", "bbb", "ccc"}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}

	src = linq.FromSlice([]string{})
	r, err = linq.ToSlice(
		linq.DefaultIfEmpty(src, "default"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp = []string{"default"}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
