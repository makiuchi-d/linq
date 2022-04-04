package linq

import "golang.org/x/exp/constraints"

// Sum calculates the sum of a integer sequence.
func Sum[T constraints.Integer](src Enumerator[T]) (int, error) {
	sum := 0
	err := ForEach(src, func(v T) error {
		sum += int(v)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return sum, nil
}

// Sumf calculates the sum of a floating number sequence.
func Sumf[T constraints.Float](src Enumerator[T]) (float64, error) {
	var sum float64
	err := ForEach(src, func(v T) error {
		sum += float64(v)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return sum, nil
}

// SumByFunc computes the sum of the sequence of values that are obtained by invoking a transform function on each element of the input sequence.
func SumByFunc[T any, K constraints.Integer](src Enumerator[T], selector func(v T) (K, error)) (int, error) {
	return Sum(Select(src, selector))
}

// SumByFuncf computes the sum of the sequence of values that are obtained by invoking a transform function on each element of the input sequence.
func SumByFuncf[T any, K constraints.Float](src Enumerator[T], selector func(v T) (K, error)) (float64, error) {
	return Sumf(Select(src, selector))
}
