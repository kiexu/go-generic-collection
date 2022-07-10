package gcollection

type Set[T comparable] interface {
	Collection

	// Contains check if element exists
	Contains(T) bool

	// Add an element to set
	// return:  true: if not exist before
	//			false: if already exist or volume reaches maxVolume
	Add(T) bool

	// Remove an element
	// return:  true: exist and success
	//			false: not exist before
	Remove(T) bool

	// ToSlice Get slice of elements
	ToSlice() []T
}
