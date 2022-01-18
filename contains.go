package linq

// Contains cetermines whether a sequence contains a specified element.
func Contains[T comparable](src Enumerator[T], val T) (bool, error) {
	for {
		t, err := src.Next()
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
func ContainsFunc[T any](src Enumerator[T], val T, equals func(T, T) (bool, error)) (bool, error) {
	for {
		t, err := src.Next()
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
