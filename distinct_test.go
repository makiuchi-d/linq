package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestDistinct(t *testing.T) {
	ages := []int{21, 46, 46, 55, 17, 21, 55, 55, 25, 25}

	e := linq.Distinct(
		linq.FromSlice(ages),
		func(a, b int) (bool, error) { return a == b, nil },
		func(a int) (int, error) { return a / 10, nil })

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []int{21, 46, 55, 17, 25}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
