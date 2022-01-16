package linq

// All determines whether all elements of a sequence satisfy a condition.
func All[T any](src Enumerator[T], cond func(v T) (bool, error)) (bool, error) {
	for {
		v, err := src.Next()
		if err != nil {
			if isEOC(err) {
				return true, nil
			}
			return false, err
		}
		ok, err := cond(v)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
}
