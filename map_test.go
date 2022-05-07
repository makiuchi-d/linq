package linq_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

func TestMaps(t *testing.T) {
	m := map[int]string{
		1:  "One",
		3:  "Three",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		20: "Twenty",
	}

	e1 := linq.FromMap(m)

	e2 := linq.Where(e1, func(kv linq.KeyValue[int, string]) (bool, error) { return kv.Key%2 != 0, nil })
	e3 := linq.Select(e2, func(kv linq.KeyValue[int, string]) (string, error) { return kv.Value, nil })
	s, err := linq.ToSlice(e3)

	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := []string{"One", "Three", "Eleven"}
	sort.Strings(s)
	sort.Strings(exp)
	if !reflect.DeepEqual(s, exp) {
		t.Fatalf("wants: %#v, got %#v", exp, s)
	}
}

func TestSliceToMap(t *testing.T) {
	s := []byte{5, 10, 15, 20}

	e1 := linq.FromSlice(s)
	e2 := linq.Select(e1, func(v byte) (linq.KeyValue[byte, string], error) {
		return linq.KeyValue[byte, string]{v, fmt.Sprintf("%02x", v)}, nil
	})
	m, err := linq.ToMap(e2)
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := map[byte]string{
		5:  "05",
		10: "0a",
		15: "0f",
		20: "14",
	}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("wants: %#v, got %#v", exp, m)
	}
}

func TestToMapFunc(t *testing.T) {
	s := []byte{5, 10, 15, 20}
	e1 := linq.FromSlice(s)
	m, err := linq.ToMapFunc(e1, func(v byte) (int, string, error) {
		return int(v), fmt.Sprintf("%02x", v), nil
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := map[int]string{
		5:  "05",
		10: "0a",
		15: "0f",
		20: "14",
	}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("wants: %#v, got %#v", exp, m)
	}
}
