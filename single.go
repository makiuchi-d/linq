package linq

// Single returns the only element of a sequence that satisfies a specified condition, and return an error InvalidOperation if more than one such element exists.
func Single[T any](src Enumerator[T], pred func(T) (bool, error)) (def T, _ error) {
	e := Where(src, pred)
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

// SingleOrDefault returns the only element of a sequence that satisfies a specified condition, or a specified default value if no such element exists; this function returns an error InvalidOperation if more than one element satisfies the condition.
func SingleOrDefault[T any](src Enumerator[T], pred func(T) (bool, error), defaultValue T) (def T, _ error) {
	e := Where(src, pred)
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
