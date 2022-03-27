package linq

type emptyEnumerator[T any] struct{}

// Empty returns an empty IEnumerable<T> that has the specified type argument.
func Empty[T any]() Enumerator[T] {
	return emptyEnumerator[T]{}
}

func (emptyEnumerator[T]) Next() (def T, _ error) {
	return def, EOC
}
