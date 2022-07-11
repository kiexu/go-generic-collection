package gcollection

type Iterator[T any] interface {

	// HasNext whether next element exists
	HasNext() bool

	// Next get one element
	Next() T
}
