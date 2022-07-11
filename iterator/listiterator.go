package iterator

import (
	"container/list"
	"github.com/kiexu/go-generic-collection"
)

var _ gcollection.Iterator[struct{}] = new(ListIterator[struct{}])

// ListIterator beware of type of list elements
type ListIterator[T any] struct {
	ptr  *list.Element
	list *list.List
}

func NewListIterator[T any](list *list.List) gcollection.Iterator[T] {
	return &ListIterator[T]{
		ptr:  nil,
		list: list,
	}
}

func (s *ListIterator[T]) HasNext() bool {
	if s.ptr == nil {
		return s.list != nil && s.list.Len() > 0
	}
	return s.ptr.Next() != nil

}

func (s *ListIterator[T]) Next() T {
	if s.ptr == nil {
		s.ptr = s.list.Front()
	} else {
		s.ptr = s.ptr.Next()
	}
	return s.ptr.Value.(T)
}

func (s *ListIterator[T]) Ptr() *list.Element {
	return s.ptr
}

func (s *ListIterator[T]) SetPtr(ptr *list.Element) {
	s.ptr = ptr
}

func (s *ListIterator[T]) List() *list.List {
	return s.list
}

func (s *ListIterator[T]) SetList(list *list.List) {
	s.list = list
}
