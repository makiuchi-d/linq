package linq

// First returns the first element in a sequence
func First[T any, E IEnumerable[T]](src E) (def T, _ error) {
	e := src()
	v, err := e.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	return v, nil
}

// FirstOrDefault returns the first element of the sequence, or a specified default value if no such element is found.
func FirstOrDefault[T any, E IEnumerable[T]](src E, defaultValue T) (T, error) {
	v, err := First(src)
	if err != nil {
		if isInvalidOperation(err) {
			err = nil
		}
		return defaultValue, err
	}
	return v, nil
}
