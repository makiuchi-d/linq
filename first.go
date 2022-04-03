package linq

// First returns the first element in a sequence that satisfies a specified condition.
func First[T any](src Enumerator[T], pred func(T) (bool, error)) (def T, _ error) {
	e := Where(src, pred)
	v, err := e.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	return v, nil
}

// FirstOrDefault returns the first element of the sequence that satisfies a condition, or a specified default value if no such element is found.
func FirstOrDefault[T any](src Enumerator[T], pred func(T) (bool, error), defaultValue T) (T, error) {
	v, err := First(src, pred)
	if err != nil {
		if isInvalidOperation(err) {
			err = nil
		}
		return defaultValue, err
	}
	return v, nil
}
