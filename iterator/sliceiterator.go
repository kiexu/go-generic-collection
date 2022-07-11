package iterator

import (
	"github.com/kiexu/go-generic-collection"
)

var _ gcollection.Iterator[struct{}] = new(SliceIterator[struct{}])

type SliceIterator[T any] struct {
	idx  int
	list []T
}

func NewSliceIterator[T any](list []T) gcollection.Iterator[T] {
	return &SliceIterator[T]{
		idx:  -1,
		list: list,
	}
}

func (s *SliceIterator[T]) HasNext() bool {
	return s.idx+1 < len(s.list)
}

func (s *SliceIterator[T]) Next() T {
	s.idx += 1
	return s.list[s.idx]
}

func (s *SliceIterator[T]) Idx() int {
	return s.idx
}

func (s *SliceIterator[T]) SetIdx(idx int) {
	s.idx = idx
}

func (s *SliceIterator[T]) List() []T {
	return s.list
}

func (s *SliceIterator[T]) SetList(list []T) {
	s.list = list
}
