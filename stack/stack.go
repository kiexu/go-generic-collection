package stack

import (
	"container/list"
	"github.com/kiexu/go-generic-collection"
	"github.com/kiexu/go-generic-collection/iterator"
)

var _ gcollection.Stack[struct{}] = new(Stack[struct{}])

type Stack[T comparable] struct {
	eList *list.List
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		eList: list.New(),
	}
}

// Size get current size
func (s *Stack[T]) Size() int {
	return s.eList.Len()
}

// IsEmpty check is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Push one element
// return: element itself
func (s *Stack[T]) Push(t T) T {
	s.eList.PushBack(t)
	return t
}

// Pop remove and return top of the stack
// return: whether actually removed
func (s *Stack[T]) Pop() (res T, ok bool) {
	if s.IsEmpty() {
		return
	}
	return s.eList.Remove(s.eList.Back()).(T), true
}

// Peek return top of the stack without removing
// return: whether actually exist
func (s *Stack[T]) Peek() (res T, ok bool) {
	if s.IsEmpty() {
		return
	}
	return s.eList.Back().Value.(T), true
}

// Clear all the elements
func (s *Stack[T]) Clear() {
	s.eList = list.New()
}

// Iterator get an Iterator
func (s *Stack[T]) Iterator() gcollection.Iterator[T] {
	return iterator.NewListIterator[T](s.eList)
}

// ForEach run ConsumeFunc on each element
func (s *Stack[T]) ForEach(c gcollection.ConsumeFunc[T]) {
	for ele := s.eList.Front(); ele != nil; ele = ele.Next() {
		c(ele.Value.(T))
	}
}

// ToSlice Get slice of elements from bottom to top
func (s *Stack[T]) ToSlice() []T {
	chain := s.eList
	res := make([]T, 0, chain.Len())

	for ele := chain.Front(); ele != nil; ele = ele.Next() {
		res = append(res, ele.Value.(T))
	}
	return res
}
