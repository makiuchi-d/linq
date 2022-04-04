package linq

import "golang.org/x/exp/constraints"

// Min returns the minimum value in a sequence of values.
func Min[T constraints.Ordered](src Enumerator[T]) (def T, _ error) {
	min, err := src.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	for {
		v, err := src.Next()
		if err != nil {
			if isEOC(err) {
				return min, nil
			}
			return def, err
		}
		if err != nil {
			return def, err
		}
		if v < min {
			min = v
		}
	}
}

// MinBy returns the minimum value in a generic sequence according to a specified key selector function.
func MinBy[T any, K constraints.Ordered](src Enumerator[T], keySelector func(T) (K, error)) (def T, _ error) {
	min, err := src.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	mink, err := keySelector(min)
	if err != nil {
		return def, err
	}
	for {
		v, err := src.Next()
		if err != nil {
			if isEOC(err) {
				return min, nil
			}
			return def, err
		}
		k, err := keySelector(v)
		if err != nil {
			return def, err
		}
		if k < mink {
			min, mink = v, k
		}
	}
}

// MinByFunc returns the minimum value in a generic sequence according to a comparer function.
func MinByFunc[T any](src Enumerator[T], less func(a, b T) (bool, error)) (def T, _ error) {
	min, err := src.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	for {
		v, err := src.Next()
		if err != nil {
			if isEOC(err) {
				return min, nil
			}
			return def, err
		}
		b, err := less(v, min)
		if err != nil {
			return def, err
		}
		if b {
			min = v
		}
	}
}
