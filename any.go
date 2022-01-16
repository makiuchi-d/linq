package linq

// Any determines whether any element of a sequence satisfies a condition.
func Any[T any](src Enumerator[T], pred func(v T) (bool, error)) (bool, error) {
	for {
		v, err := src.Next()
		if err != nil {
			if isEOC(err) {
				return false, nil
			}
			return false, err
		}
		ok, err := pred(v)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
}
