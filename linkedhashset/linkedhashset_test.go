package linkedhashset

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestHashSet(t *testing.T) {

	s := NewLinkedHashSet[string]()
	loop := 99

	assert.Equal(t, s.Add("1"), true)
	assert.Equal(t, s.Add("2"), true)
	assert.Equal(t, s.Add("3"), true)
	assert.Equal(t, s.Add("1"), false)
	for i := 0; i < loop; i++ {
		assert.DeepEqual(t, s.ToSlice(), []string{"1", "2", "3"})
	}
	assert.Equal(t, s.Size(), 3)
	assert.Equal(t, s.Contains("2"), true)
	assert.Equal(t, s.Remove("4"), false)
	assert.Equal(t, s.Remove("3"), true)
	assert.Equal(t, s.Contains("3"), false)
	for i := 0; i < loop; i++ {
		assert.DeepEqual(t, s.ToSlice(), []string{"1", "2"})
	}
	assert.Equal(t, s.IsEmpty(), false)
	s.Clear()
	assert.Equal(t, s.Size(), 0)
	assert.Equal(t, s.IsEmpty(), true)
}
