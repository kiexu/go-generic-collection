package gcollection

type Collection interface {

	// Size get current size
	Size() int

	// IsEmpty check is empty
	IsEmpty() bool

	// Clear all elements
	Clear()
}
