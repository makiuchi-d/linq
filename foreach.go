package linq

// ForEach performs the specified function on each element of the specified enumerator.
func ForEach[T any, E IEnumerable[T]](src E, f func(T) error) error {
	e := src()
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return nil
			}
			return err
		}
		err = f(v)
		if err != nil {
			return err
		}
	}
}
