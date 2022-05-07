package linq_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestJoin(t *testing.T) {
	nums := linq.Range(3, 6)
	strs := linq.FromSlice(
		[]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"})

	e := linq.Join(nums, strs,
		func(n int) (int, error) { return n, nil },
		func(s string) (int, error) { return len(s), nil },
		func(n int, s string) (string, error) { return fmt.Sprintf("%d:%s", n, s), nil })

	r, err := linq.ToSlice(e)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []string{
		"3:one", "3:two", "3:six", "3:ten",
		"4:four", "4:five", "4:nine",
		"5:three", "5:seven", "5:eight",
	}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
