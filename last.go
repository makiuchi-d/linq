package linq

// Last returns the last element of a sequence that satisfies a specified condition.
func Last[T any, E IEnumerable[T]](src E) (def T, _ error) {
	var last T
	var eocerr error = InvalidOperation
	e := src()
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return last, eocerr
			}
			return def, err
		}
		last = v
		eocerr = nil
	}
}

// LastOrDefault returns the last element of a sequence that satisfies a condition, or a specified default value if no such element is found.
func LastOrDefault[T any, E IEnumerable[T]](src E, defaultValue T) (T, error) {
	v, err := Last(src)
	if isInvalidOperation(err) {
		return defaultValue, nil
	}
	return v, err
}
