package linq

// Single returns the only element of a sequence, and return an error InvalidOperation if more than one such element exists.
func Single[T any, E IEnumerable[T]](src E) (def T, _ error) {
	e := src()
	v, err := e.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	_, err = e.Next()
	if err == nil {
		return def, InvalidOperation
	}
	if !isEOC(err) {
		return def, err
	}
	return v, nil
}

// SingleOrDefault returns the only element of a sequence, or a specified default value if no such element exists; this function returns an error InvalidOperation if more than one element satisfies the condition.
func SingleOrDefault[T any, E IEnumerable[T]](src E, defaultValue T) (def T, _ error) {
	e := src()
	v, err := e.Next()
	if err != nil {
		if isEOC(err) {
			return defaultValue, nil
		}
		return def, err
	}
	_, err = e.Next()
	if err == nil {
		return def, InvalidOperation
	}
	if !isEOC(err) {
		return def, err
	}
	return v, nil
}
