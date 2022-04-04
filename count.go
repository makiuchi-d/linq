package linq

// Count returns a number that represents how many elements in the specified sequence satisfy a condition.
func Count[T any](src Enumerator[T], pred func(T) (bool, error)) (int, error) {
	c := 0
	err := ForEach(Where(src, pred), func(v T) error {
		c++
		return nil
	})
	if err != nil {
		return 0, err
	}
	return c, nil
}
