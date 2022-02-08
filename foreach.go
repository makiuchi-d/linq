package linq

// ForEach performs the specified action on each element of the specified enumerator.
func ForEach[T any](e Enumerator[T], f func(T) error) error {
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
