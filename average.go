package linq

import "golang.org/x/exp/constraints"

type Real interface {
	constraints.Integer | constraints.Float
}

// Average computes the average of a sequence of numeric (real number) values.
func Average[T Real, E IEnumerable[T]](src E) (float64, error) {
	n := 0
	t := float64(0)
	err := ForEach(src, func(v T) error {
		t += float64(v)
		n++
		return nil
	})
	if err != nil {
		return 0, err
	}
	if n == 0 {
		return 0, InvalidOperation
	}
	return t / float64(n), nil
}
