package linkedhashmap

import (
	"github.com/kiexu/go-generic-collection"
	"gotest.tools/v3/assert"
	"testing"
)

func TestLinkedHashMap(t *testing.T) {
	type tp struct {
		N int
	}
	var zeroV int
	var r1 int
	var r2 bool
	loop := 99

	s := NewLinkedHashMap[tp, int]()

	r1, r2 = s.Put(tp{1}, 1)
	assert.Equal(t, r1, zeroV)
	assert.Equal(t, r2, false)

	r1, r2 = s.Put(tp{2}, 2)
	assert.Equal(t, r1, zeroV)
	assert.Equal(t, r2, false)

	r1, r2 = s.Put(tp{3}, 3)
	assert.Equal(t, r1, zeroV)
	assert.Equal(t, r2, false)

	assert.Equal(t, s.Size(), 3)

	for i := 0; i < loop; i++ {

		assert.DeepEqual(t, s.Keys(), []tp{{1}, {2}, {3}})
		assert.DeepEqual(t, s.Values(), []int{1, 2, 3})

		var entry1 gcollection.MapEntry[tp, int] = &LinkedHashEntry[tp, int]{key: tp{1}, value: 1}
		var entry2 gcollection.MapEntry[tp, int] = &LinkedHashEntry[tp, int]{key: tp{2}, value: 2}
		var entry3 gcollection.MapEntry[tp, int] = &LinkedHashEntry[tp, int]{key: tp{3}, value: 3}
		assert.Equal(t, s.Entries()[0].Value(), entry1.Value())
		assert.Equal(t, s.Entries()[0].Key(), entry1.Key())
		assert.Equal(t, s.Entries()[1].Value(), entry2.Value())
		assert.Equal(t, s.Entries()[1].Key(), entry2.Key())
		assert.Equal(t, s.Entries()[2].Value(), entry3.Value())
		assert.Equal(t, s.Entries()[2].Key(), entry3.Key())
	}

	r1, r2 = s.Get(tp{2})
	assert.Equal(t, r1, 2)
	assert.Equal(t, r2, true)

	r1, r2 = s.Get(tp{4})
	assert.Equal(t, r1, zeroV)
	assert.Equal(t, r2, false)

	assert.Equal(t, s.Size(), 3)

	s.Remove(tp{3})
	assert.Equal(t, s.Size(), 2)

	r1, r2 = s.Get(tp{2})
	assert.Equal(t, r1, 2)
	assert.Equal(t, r2, true)

	r1, r2 = s.Get(tp{3})
	assert.Equal(t, r1, zeroV)
	assert.Equal(t, r2, false)

	s.Clear()
	assert.Equal(t, s.Size(), 0)

	r1, r2 = s.Put(tp{1}, 1)
	assert.Equal(t, r1, zeroV)
	assert.Equal(t, r2, false)
}

func TestLinkedHashMap_ForEach(t *testing.T) {

	s := NewLinkedHashMap[string, int]()
	s.Put("1", 1)
	s.Put("2", 2)
	s.Put("3", 3)

	str := ""
	sum := 0
	s.ForEach(func(e gcollection.MapEntry[string, int]) {
		str += e.Key()
		sum += e.Value()
	})

	assert.Equal(t, str, "123")
	assert.Equal(t, sum, 6)
}

func TestLinkedHashMap_Iterator(t *testing.T) {

	s := NewLinkedHashMap[string, int]()
	s.Put("1", 1)
	s.Put("2", 2)
	s.Put("3", 3)

	str := ""
	sum := 0
	itr := s.Iterator()
	for itr.HasNext() {
		next := itr.Next()
		str += next.Key()
		sum += next.Value()
	}

	assert.Equal(t, str, "123")
	assert.Equal(t, sum, 6)
}
