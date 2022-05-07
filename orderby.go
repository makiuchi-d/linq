package linq

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type orderedEnumerator[T any] struct {
	src     Enumerator[T]
	newcmps []func([]T) (comparer, error)
	sorted  []T
	i       int
}

// OrderBy sorts the elements of a sequence in ascending order according to a key.
func OrderBy[T any, K constraints.Ordered, E IEnumerable[T]](src E, keySelector func(T) (K, error)) OrderedEnumerable[T] {
	return func() Enumerator[T] {
		return &orderedEnumerator[T]{
			src: src(),
			newcmps: []func([]T) (comparer, error){
				newKeyComparer[kCmpAsc[K]](keySelector),
			},
		}
	}
}

// OrderByDescending sorts the elements of a sequence in descending order according to a key.
func OrderByDescending[T any, K constraints.Ordered, E IEnumerable[T]](src E, keySelector func(T) (K, error)) OrderedEnumerable[T] {
	return func() Enumerator[T] {
		return &orderedEnumerator[T]{
			src: src(),
			newcmps: []func([]T) (comparer, error){
				newKeyComparer[kCmpDesc[K]](keySelector),
			},
		}
	}
}

// ThenBy performs a subsequent ordering of the elements in a sequence in ascending order according to a key.
func ThenBy[T any, K constraints.Ordered](src OrderedEnumerable[T], keySelector func(T) (K, error)) OrderedEnumerable[T] {
	return func() Enumerator[T] {
		oe := src().(*orderedEnumerator[T])
		return &orderedEnumerator[T]{
			src:     oe.src,
			newcmps: append(oe.newcmps, newKeyComparer[kCmpAsc[K]](keySelector)),
		}
	}
}

// ThenByDescending performs a subsequent ordering of the elements in a sequence in descending order, according to a key.
func ThenByDescending[T any, K constraints.Ordered](src OrderedEnumerable[T], keySelector func(T) (K, error)) OrderedEnumerable[T] {
	return func() Enumerator[T] {
		oe := src().(*orderedEnumerator[T])
		return &orderedEnumerator[T]{
			src:     oe.src,
			newcmps: append(oe.newcmps, newKeyComparer[kCmpDesc[K]](keySelector)),
		}
	}
}

func (e *orderedEnumerator[T]) Next() (def T, _ error) {
	if e.sorted == nil {
		s, err := doSort(e.src, e.newcmps)
		if err != nil {
			return def, err
		}
		e.sorted = s
	}
	if e.i >= len(e.sorted) {
		return def, EOC
	}
	i := e.i
	e.i++
	return e.sorted[i], nil
}

type comparer interface {
	compare(i, j int) int
	swap(i, j int)
}

type kCmpAsc[K constraints.Ordered] []K
type kCmpDesc[K constraints.Ordered] []K

var _ comparer = kCmpAsc[int]{}
var _ comparer = kCmpDesc[int]{}

type keyComparer[K constraints.Ordered] interface {
	kCmpAsc[K] | kCmpDesc[K]
	comparer
}

func newKeyComparer[C keyComparer[K], T any, K constraints.Ordered](keysel func(T) (K, error)) func(s []T) (comparer, error) {
	return func(s []T) (comparer, error) {
		ks := make([]K, len(s))
		for i, t := range s {
			k, err := keysel(t)
			if err != nil {
				return nil, err
			}
			ks[i] = k
		}
		return C(ks), nil
	}
}

func (c kCmpAsc[K]) compare(i, j int) int {
	switch {
	case c[i] < c[j]:
		return -1
	case c[i] > c[j]:
		return 1
	default:
		return 0
	}
}

func (c kCmpAsc[K]) swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c kCmpDesc[K]) compare(i, j int) int {
	switch {
	case c[i] < c[j]:
		return 1
	case c[i] > c[j]:
		return -1
	default:
		return 0
	}
}

func (c kCmpDesc[K]) swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func doSort[T any](src Enumerator[T], newcmps []func([]T) (comparer, error)) ([]T, error) {
	s, err := toSlice(src)
	if err != nil {
		return nil, err
	}
	cmps := make([]comparer, len(newcmps))
	for i, newcmp := range newcmps {
		cmps[i], err = newcmp(s)
		if err != nil {
			return nil, err
		}
	}
	sort.Sort(&sorter[T]{src: s, cmps: cmps})
	return s, nil
}

type sorter[T any] struct {
	src  []T
	cmps []comparer
}

var _ sort.Interface = &sorter[int]{}

func (s *sorter[T]) Len() int {
	return len(s.src)
}

func (s *sorter[T]) Less(i, j int) bool {
	for _, cmp := range s.cmps {
		switch c := cmp.compare(i, j); true {
		case c < 0:
			return true
		case c > 0:
			return false
		}
	}
	return true
}

func (s *sorter[T]) Swap(i, j int) {
	s.src[i], s.src[j] = s.src[j], s.src[i]
	for _, cmp := range s.cmps {
		cmp.swap(i, j)
	}
}
