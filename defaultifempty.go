package linq

type defaultIfEmptyEnumerator[T any] struct {
	src  Enumerator[T]
	def  T
	rest Enumerator[T]
}

// DefaultIfEmpty returns the elements of the specified sequence or the specified value in a singleton collection if the sequence is empty.
func DefaultIfEmpty[T any, E IEnumerable[T]](src E, defaultValue T) Enumerable[T] {
	return func() Enumerator[T] {
		return &defaultIfEmptyEnumerator[T]{src: src(), def: defaultValue}
	}
}

func (e *defaultIfEmptyEnumerator[T]) Next() (def T, _ error) {
	if e.rest != nil {
		return e.rest.Next()
	}
	v, err := e.src.Next()
	if err != nil {
		if isEOC(err) {
			e.rest = Empty[T]()()
			return e.def, nil
		}
		return def, err
	}
	e.rest = e.src
	return v, nil
}
