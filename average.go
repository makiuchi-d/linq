package linq

import "golang.org/x/exp/constraints"

type Real interface {
	constraints.Integer | constraints.Float
}

// Average computes the average of a sequence of numeric (real number) values.
func Average[T Real](src Enumerator[T]) (float64, error) {
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

// AverageFunc computes the average of a sequence of numeric values that are obtained by invoking a transform function on each element of the input sequence.
func AverageFunc[T any, U Real](src Enumerator[T], selector func(T) (U, error)) (float64, error) {
	return Average(Select(src, selector))
}
