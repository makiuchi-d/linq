// LINQ
package linq

// IEnumerable[T] is a queryable collection.
type IEnumerable[T any] interface {
	Enumerable[T] | OrderedEnumerable[T]
}

// Enumerable[T] is an implementation of IEnumerable[T].
type Enumerable[T any] func() (Enumerator[T], error)

// OrderedEnumerable[T] is an implementation of IEnumerable[T], which is generated by OrderBy or ThenBy.
type OrderedEnumerable[T any] func() (Enumerator[T], error)

// Enumerator[T] is an enumerator of the collection.
type Enumerator[T any] interface {
	// Next returns a next element of this collection.
	// It returns EOC as an `error` when it reaches the end of the collection.
	Next() (T, error)
}

// Error : LINQ error type.
type Error string

const (
	// EOC : End of the collection.
	EOC Error = "End of the collection"

	// OutOfRange : Index out of range.
	OutOfRange Error = "Out of range"

	// InvalidOperation : Invalid operation such as no element satisfying the condition.
	InvalidOperation Error = "Invalid operation"
)

func (e Error) Error() string {
	return string(e)
}
