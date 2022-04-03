package linq

// First returns the first element in a sequence that satisfies a specified condition.
func First[T any](src Enumerator[T], pred func(T) (bool, error)) (def T, _ error) {
	for {
		v, err := src.Next()
		if err != nil {
			if isEOC(err) {
				err = InvalidOperation
			}
			return def, err
		}
		b, err := pred(v)
		if err != nil {
			return def, err
		}
		if b {
			return v, nil
		}
	}
}

// FirstOrDefault returns the first element of the sequence that satisfies a condition, or a specified default value if no such element is found.
func FirstOrDefault[T any](src Enumerator[T], pred func(T) (bool, error), def T) (T, error) {
	v, err := First(src, pred)
	if err != nil {
		if isInvalidOperation(err) {
			err = nil
		}
		return def, err
	}
	return v, nil
}
