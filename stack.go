package gcollection

type Stack[T comparable] interface {
	Collection
	Iterable[T]

	// Push one element
	// return: element itself
	Push(T) T

	// Pop remove and return top of the stack
	// return: whether actually removed
	Pop() (T, bool)

	// Peek return top of the stack without removing
	// return: whether actually exist
	Peek() (T, bool)
}
