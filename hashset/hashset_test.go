package hashset

import (
	"gotest.tools/v3/assert"
	"testing"
)

// TestHashSet https://pkg.go.dev/gotest.tools/v3
func TestHashSet(t *testing.T) {

	s := NewHashSet[string]()
	assert.Equal(t, s.Add("1"), true)
	assert.Equal(t, s.Add("2"), true)
	assert.Equal(t, s.Add("3"), true)
	assert.Equal(t, s.Add("1"), false)
	assert.Equal(t, s.Size(), 3)
	assert.Equal(t, s.Contains("2"), true)
	assert.Equal(t, s.Remove("4"), false)
	assert.Equal(t, s.Remove("3"), true)
	assert.Equal(t, s.Contains("3"), false)
	assert.Equal(t, len(s.ToSlice()), 2)
	assert.Equal(t, s.IsEmpty(), false)
	s.Clear()
	assert.Equal(t, s.Size(), 0)
	assert.Equal(t, s.IsEmpty(), true)
}

func TestHashSet_Iterator(t *testing.T) {

	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	sum := 0
	itr := s.Iterator()
	for itr.HasNext() {
		sum += itr.Next()
	}

	assert.Equal(t, sum, 6)
}

func TestHashSet_ForEach(t *testing.T) {

	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	sum := 0
	s.ForEach(func(a int) {
		sum += a
	})

	assert.Equal(t, sum, 6)
}
