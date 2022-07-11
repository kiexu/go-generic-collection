package linkedhashset

import (
	"container/list"
	"github.com/kiexu/go-generic-collection"
	"github.com/kiexu/go-generic-collection/iterator"
)

var _ gcollection.Set[struct{}] = new(LinkedHashSet[struct{}])

// LinkedHashSet
// not thread-safe
type LinkedHashSet[T comparable] struct {
	eList *list.List          // list.Element contains element
	eMap  map[T]*list.Element // key: element, value: element in Element
}

func NewLinkedHashSet[T comparable]() *LinkedHashSet[T] {
	return &LinkedHashSet[T]{
		eList: list.New(),
		eMap:  make(map[T]*list.Element),
	}
}

// Size get current size
func (s *LinkedHashSet[T]) Size() int {
	return s.eList.Len()
}

// IsEmpty check if is empty
func (s *LinkedHashSet[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Contains check if element exists
func (s *LinkedHashSet[T]) Contains(e T) bool {
	_, exist := s.eMap[e]
	return exist
}

// Add an element to set
// return:  true if not exist before
//			false if already exist or volume reaches maxVolume
func (s *LinkedHashSet[T]) Add(e T) bool {

	if _, exist := s.eMap[e]; exist {
		return false
	}
	s.eMap[e] = s.eList.PushBack(e)

	return true
}

// Remove an element
// return:  true: exist and success
//			false: not exist before
func (s *LinkedHashSet[T]) Remove(e T) bool {

	elePtr, exist := s.eMap[e]
	if !exist {
		return false
	}

	s.eList.Remove(elePtr)
	delete(s.eMap, e)

	return true
}

// Clear all the elements
// maxVolume will not modify
func (s *LinkedHashSet[T]) Clear() {

	s.eList = list.New()
	s.eMap = make(map[T]*list.Element)

	return
}

// Iterator get an Iterator
func (s *LinkedHashSet[T]) Iterator() gcollection.Iterator[T] {
	return iterator.NewListIterator[T](s.eList)
}

// ForEach run ConsumeFunc on each element
func (s *LinkedHashSet[T]) ForEach(c gcollection.ConsumeFunc[T]) {
	for ele := s.eList.Front(); ele != nil; ele = ele.Next() {
		c(ele.Value.(T))
	}
}

// ToSlice Get add-base-ordered slice of objects
func (s *LinkedHashSet[T]) ToSlice() []T {

	chain := s.eList
	res := make([]T, 0, chain.Len())

	for ele := chain.Front(); ele != nil; ele = ele.Next() {
		res = append(res, ele.Value.(T))
	}
	return res
}
