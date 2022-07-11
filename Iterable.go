package gcollection

type Iterable[T any] interface {

	// Iterator get an Iterator
	Iterator() Iterator[T]

	// ForEach run ConsumeFunc on each element
	ForEach(ConsumeFunc[T])

	// ToSlice Get slice of elements
	ToSlice() []T
}
