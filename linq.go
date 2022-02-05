// LINQ
package linq

// Enumerator[T] is a queryable collection
type Enumerator[T any] interface {

	// Next returns a next element of this collection.
	// It returns EOC as an `error` when it reaches the end of the collection.
	Next() (T, error)

	// This package does not support Reset() because many of the IEnumerator<T>
	// used in the C# LINQ method did not support the Reset method.
	// Reset() error
}

// Error : LINQ error type.
type Error string

// EOC : End of the collection
const EOC Error = "End of the collection"

func (e Error) Error() string {
	return string(e)
}
