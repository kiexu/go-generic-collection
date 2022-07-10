package stack

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestStack(t *testing.T) {

	s := NewStack[int]()
	var r1 int
	var r2 bool
	var zero int

	assert.Equal(t, s.Size(), 0)
	assert.Equal(t, s.Push(1), 1)
	assert.Equal(t, s.Push(2), 2)
	assert.Equal(t, s.Push(3), 3)

	assert.DeepEqual(t, s.ToSlice(), []int{1, 2, 3})

	r1, r2 = s.Pop()
	assert.Equal(t, r1, 3)
	assert.Equal(t, r2, true)

	assert.Equal(t, s.Push(4), 4)

	assert.DeepEqual(t, s.ToSlice(), []int{1, 2, 4})

	r1, r2 = s.Pop()
	assert.Equal(t, r1, 4)
	assert.Equal(t, r2, true)

	r1, r2 = s.Pop()
	assert.Equal(t, r1, 2)
	assert.Equal(t, r2, true)

	assert.Equal(t, s.Size(), 1)

	r1, r2 = s.Pop()
	assert.Equal(t, r1, 1)
	assert.Equal(t, r2, true)

	r1, r2 = s.Pop()
	assert.Equal(t, r1, zero)
	assert.Equal(t, r2, false)

	r1, r2 = s.Pop()
	assert.Equal(t, r1, zero)
	assert.Equal(t, r2, false)

	assert.Equal(t, s.Size(), 0)
}
