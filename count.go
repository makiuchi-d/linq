package linq

// Count returns a number that represents how many elements in the specified sequence
func Count[T any, E IEnumerable[T]](src E) (int, error) {
	c := 0
	err := ForEach(src, func(v T) error {
		c++
		return nil
	})
	if err != nil {
		return 0, err
	}
	return c, nil
}
