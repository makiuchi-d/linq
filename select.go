package linq

type selectEnumerator[T, U any] struct {
	src Enumerator[T]
	sel func(v T) (U, error)
}

// Select projects each element of a sequence into a new form.
func Select[T, U any](src Enumerator[T], selector func(v T) (U, error)) Enumerator[U] {
	return &selectEnumerator[T, U]{src: src, sel: selector}
}

func (e *selectEnumerator[T, U]) Next() (U, error) {
	v, err := e.src.Next()
	if err != nil {
		var u U
		return u, err
	}

	return e.sel(v)
}
