package gcollection

type Map[K, V comparable] interface {
	Collection
	Iterable[MapEntry[K, V]]

	// ContainsKey check if key exists
	ContainsKey(key K) bool

	// ContainsValue check if value exists
	ContainsValue(value V) bool

	// Put a key-value and try return previous value of the given key
	// return: previous: previous value
	// 		   exist: if previous value exists
	Put(key K, value V) (V, bool)

	// Get value by key
	// return: value: value of given key
	// 		   exist: if value exists
	Get(key K) (V, bool)

	// Remove key-value
	// return: success or not exist
	Remove(key K) bool

	// Keys get all keys
	Keys() []K

	// Values get all values
	Values() []V
}

// MapEntry common map entry
type MapEntry[K, V comparable] interface {
	Key() K
	Value() V
	SetKey(key K)
	SetValue(value V)
}
