package linq_test

import (
	"reflect"
	"testing"

	"github.com/makiuchi-d/linq"
)

func TestGroupJoin(t *testing.T) {
	type Person struct {
		Name string
	}

	type Pet struct {
		Name  string
		Owner *Person
	}

	alice := &Person{"Alice"}
	bob := &Person{"Bob"}
	charlie := &Person{"Charlie"}

	abby := &Pet{"Abby", alice}
	bailey := &Pet{"Bailey", bob}
	bella := &Pet{"Bella", bob}
	cody := &Pet{"Cody", charlie}

	people := linq.FromSlice([]*Person{alice, bob, charlie})
	pets := linq.FromSlice([]*Pet{cody, bella, abby, bailey})

	r, err := linq.ToSlice(
		linq.GroupJoin(
			people, pets,
			func(p *Person) (*Person, error) { return p, nil },
			func(p *Pet) (*Person, error) { return p.Owner, nil },
			func(person *Person, pets linq.Enumerator[*Pet]) ([]string, error) {
				names, err := linq.ToSlice(
					linq.Select(pets, func(p *Pet) (string, error) { return p.Name, nil }))
				if err != nil {
					return nil, err
				}
				return append([]string{person.Name}, names...), nil
			}))
	if err != nil {
		t.Fatalf("%v", err)
	}

	exp := [][]string{
		{"Alice", "Abby"},
		{"Bob", "Bella", "Bailey"},
		{"Charlie", "Cody"},
	}
	if !reflect.DeepEqual(r, exp) {
		t.Fatalf("%v, wants %v", r, exp)
	}
}
