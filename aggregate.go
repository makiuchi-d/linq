package linq

// Aggregate applies an accumulator function over a sequence.
func Aggregate[T, U any](src Enumerator[T], seed U, fn func(acc U, v T) (U, error)) (U, error) {
	acc := seed
	err := ForEach(src, func(v T) (err error) {
		acc, err = fn(acc, v)
		return err
	})
	return acc, err
}
