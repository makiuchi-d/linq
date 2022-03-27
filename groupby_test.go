package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestGroupBy(t *testing.T) {
	src := linq.FromSlice([]string{
		"blueberry", "chimpanzee", "abacus", "banana", "apple", "cheese", "elephant", "umbrella", "anteater",
	})

	e := linq.GroupBy(src, func(s string) (byte, error) { return s[0], nil })

	mexp := map[byte][]string{
		'a': {"abacus", "anteater", "apple"},
		'b': {"banana", "blueberry"},
		'c': {"cheese", "chimpanzee"},
		'e': {"elephant"},
		'u': {"umbrella"},
	}

	linq.ForEach(e, func(grp linq.Grouping[byte, string]) error {
		s, err := linq.ToSlice[string](
			linq.OrderBy[string](grp, func(s string) (string, error) { return s, nil }))
		if err != nil {
			t.Fatalf("%c: %v", grp.Key(), err)
		}
		k := grp.Key()
		exp, ok := mexp[k]
		if !ok {
			t.Fatalf("invalid key: %c", k)
		}
		if !reflect.DeepEqual(s, exp) {
			t.Fatalf("%c: %v, wants %v", k, s, exp)
		}
		delete(mexp, k)
		return nil
	})
	if len(mexp) > 0 {
		t.Fatalf("unreturned items: %v", mexp)
	}
}
