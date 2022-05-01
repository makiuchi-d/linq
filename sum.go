package linq

import "golang.org/x/exp/constraints"

// Sum calculates the sum of a integer sequence.
func Sum[T constraints.Integer, E IEnumerable[T]](src E) (int, error) {
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
func Sumf[T constraints.Float, E IEnumerable[T]](src E) (float64, error) {
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
