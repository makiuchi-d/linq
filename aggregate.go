package linq

// Aggregate applies an accumulator function over a sequence.
func Aggregate[S, T any, E IEnumerable[S]](src E, seed T, fn func(acc T, v S) (T, error)) (T, error) {
	acc := seed
	err := ForEach(src, func(v S) (err error) {
		acc, err = fn(acc, v)
		return err
	})
	return acc, err
}
