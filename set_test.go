package gocoll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Contains(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
	assert.False(t, s.Contains(4))
}

func TestSet_Add(t *testing.T) {
	s := NewSet[int]()
	added := s.Add(1)
	assert.True(t, added)
	assert.Equal(t, 1, s.Size())
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	added = s.Add(2)
	assert.True(t, added)
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	added = s.Add(2)
	assert.False(t, added)
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	added = s.Add(3)
	assert.True(t, added)
	assert.Equal(t, 3, s.Size())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
}

func TestSet_AddAll(t *testing.T) {
	s := NewSet[int]()
	s.AddAll([]int{1, 2, 2, 3, 4, 4, 6}...)
	assert.Equal(t, 5, s.Size())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
	assert.True(t, s.Contains(4))
	assert.False(t, s.Contains(5))
	assert.True(t, s.Contains(6))
}

func TestSet_Remove(t *testing.T) {
	s := NewSet[int]()
	assert.Equal(t, 0, s.Size())
	s.Add(1)
	s.Add(2)
	s.Add(3)
	assert.Equal(t, 3, s.Size())

	s.Remove(1)
	assert.Equal(t, 2, s.Size())
	assert.False(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))

	s.Remove(2)
	assert.Equal(t, 1, s.Size())
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.True(t, s.Contains(3))

	s.Remove(3)
	assert.Equal(t, 0, s.Size())
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))
}

func TestSet_Size(t *testing.T) {
	s := NewSet[int]()
	assert.Equal(t, 0, s.Size())

	s.Add(1)
	assert.Equal(t, 1, s.Size())

	s.Add(2)
	assert.Equal(t, 2, s.Size())

	s.Add(2)
	assert.Equal(t, 2, s.Size())

	s.Add(3)
	assert.Equal(t, 3, s.Size())
}

func TestSet_Slice(t *testing.T) {
	s := NewSet[int]()

	assert.Equal(t, 0, len(s.Slice()))

	s.Add(1)
	slice := s.Slice()
	assert.Equal(t, 1, len(slice))
	assert.Equal(t, 1, slice[0])

	s.Add(2)
	slice = s.Slice()
	assert.Equal(t, 2, len(slice))
	assert.Contains(t, slice, 1, 2)

	s.Add(2)
	slice = s.Slice()
	assert.Equal(t, 2, len(slice))
	assert.Contains(t, slice, 1, 2)

	s.Add(3)
	slice = s.Slice()
	assert.Equal(t, 3, len(slice))
	assert.Contains(t, slice, 1, 2, 3)
}

func TestSet_ConstructWithItems(t *testing.T) {
	s := NewSet[int]([]int{1, 2, 2, 3, 4, 4, 6}...)
	assert.Equal(t, 5, s.Size())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
	assert.True(t, s.Contains(4))
	assert.True(t, s.Contains(6))
}

func TestSet_ZeroValue(t *testing.T) {
	var s Set[int]
	assert.False(t, s.Contains(1))
	assert.False(t, s.Remove(1))
	assert.Equal(t, 0, s.Size())
	assert.Empty(t, s.Slice())
	assert.True(t, s.Add(1))
	assert.True(t, s.Contains(1))
	assert.Equal(t, 1, s.Size())
}
