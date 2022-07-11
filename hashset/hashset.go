package hashset

import (
	"github.com/kiexu/go-generic-collection"
	"github.com/kiexu/go-generic-collection/iterator"
)

var _ gcollection.Set[struct{}] = new(HashSet[struct{}])

// HashSet
// Not thread-safe
type HashSet[T comparable] struct {
	valueMap map[T]struct{}
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		valueMap: make(map[T]struct{}),
	}
}

// Size Get current size
func (s *HashSet[T]) Size() int {
	return len(s.valueMap)
}

// IsEmpty Check if is empty
func (s *HashSet[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Contains Check if element exists
func (s *HashSet[T]) Contains(e T) bool {
	_, exist := s.valueMap[e]
	return exist
}

// Add an element to set
// return:  true if not exist before
//			false if already exist
func (s *HashSet[T]) Add(e T) bool {
	if _, exist := s.valueMap[e]; exist {
		return false
	}
	s.valueMap[e] = struct{}{}
	return true
}

// Remove an element
// return:  true: exist and success
//			false: not exist before
func (s *HashSet[T]) Remove(e T) bool {
	if _, exist := s.valueMap[e]; !exist {
		return false
	}
	delete(s.valueMap, e)
	return true
}

// Clear all the elements
func (s *HashSet[T]) Clear() {
	s.valueMap = make(map[T]struct{})
	return
}

// Iterator get an Iterator
func (s *HashSet[T]) Iterator() gcollection.Iterator[T] {
	return iterator.NewSliceIterator[T](s.ToSlice())
}

// ForEach run ConsumeFunc on each element
func (s *HashSet[T]) ForEach(c gcollection.ConsumeFunc[T]) {
	for k := range s.valueMap {
		c(k)
	}
}

// ToSlice Get disordered slice of objects
func (s *HashSet[T]) ToSlice() []T {
	res := make([]T, 0, len(s.valueMap))
	for k := range s.valueMap {
		res = append(res, k)
	}
	return res
}
