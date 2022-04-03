package linq

// Last returns the last element of a sequence that satisfies a specified condition.
func Last[T any](src Enumerator[T], pred func(T) (bool, error)) (def T, _ error) {
	var last T
	var eocerr error = InvalidOperation

	for {
		v, err := src.Next()
		if err != nil {
			if isEOC(err) {
				return last, eocerr
			}
			return def, err
		}
		b, err := pred(v)
		if err != nil {
			return def, err
		}
		if b {
			last = v
			eocerr = nil
		}
	}
}

// LastOrDefault returns the last element of a sequence that satisfies a condition, or a specified default value if no such element is found.
func LastOrDefault[T any](src Enumerator[T], pred func(T) (bool, error), defaultValue T) (T, error) {
	v, err := Last(src, pred)
	if isInvalidOperation(err) {
		return defaultValue, nil
	}
	return v, err
}
