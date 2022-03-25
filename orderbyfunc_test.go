package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestOrderByFunc(t *testing.T) {
	src := linq.FromSlice([]string{
		"grape", "passionfruit", "banana", "mango",
		"orange", "raspberry", "apple", "blueberry"})
	e := linq.OrderByFunc(src, func(a, b string) bool {
		switch l := len(a) - len(b); {
		case l < 0:
			return true
		case l > 0:
			return false
		default:
			return a < b
		}
	})
	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}
	exp := []string{
		"apple",
		"grape",
		"mango",
		"banana",
		"orange",
		"blueberry",
		"raspberry",
		"passionfruit",
	}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
