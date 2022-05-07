package linq

// Contains cetermines whether a sequence contains a specified element.
func Contains[T comparable, E IEnumerable[T]](src E, val T) (bool, error) {
	e := src()
	for {
		t, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return false, nil
			}
			return false, err
		}

		if t == val {
			return true, nil
		}
	}
}

// ContainsFunc determines whether a sequence contains a specified element by using a specified comparer function.
func ContainsFunc[T any, E IEnumerable[T]](src E, val T, equals func(T, T) (bool, error)) (bool, error) {
	e := src()
	for {
		t, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return false, nil
			}
			return false, err
		}

		eq, err := equals(t, val)
		if err != nil {
			return false, err
		}
		if eq {
			return true, nil
		}
	}
}
