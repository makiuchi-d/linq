package linq

// All determines whether all elements of a sequence satisfy a condition.
func All[T any, E IEnumerable[T]](src E, pred func(v T) (bool, error)) (bool, error) {
	e := src()
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return true, nil
			}
			return false, err
		}
		ok, err := pred(v)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
}
