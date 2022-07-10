package linkedhashmap

import (
	"gotest.tools/v3/assert"
	gcollection "kiexu/go-generic-collection"
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
