package coll

import (
	"slices"
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
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	added = s.Add(2)
	assert.True(t, added)
	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	added = s.Add(2)
	assert.False(t, added)
	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	added = s.Add(3)
	assert.True(t, added)
	assert.Equal(t, 3, s.Len())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
}

func TestSet_AddAll(t *testing.T) {
	s := NewSet[int]()
	s.AddAll([]int{1, 2, 2, 3, 4, 4, 6}...)
	assert.Equal(t, 5, s.Len())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
	assert.True(t, s.Contains(4))
	assert.False(t, s.Contains(5))
	assert.True(t, s.Contains(6))
}

func TestSet_Remove(t *testing.T) {
	s := NewSet[int]()
	assert.Equal(t, 0, s.Len())
	s.Add(1)
	s.Add(2)
	s.Add(3)
	assert.Equal(t, 3, s.Len())

	s.Remove(1)
	assert.Equal(t, 2, s.Len())
	assert.False(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))

	s.Remove(2)
	assert.Equal(t, 1, s.Len())
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.True(t, s.Contains(3))

	s.Remove(3)
	assert.Equal(t, 0, s.Len())
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))
}

func TestSet_Size(t *testing.T) {
	s := NewSet[int]()
	assert.Equal(t, 0, s.Len())

	s.Add(1)
	assert.Equal(t, 1, s.Len())

	s.Add(2)
	assert.Equal(t, 2, s.Len())

	s.Add(2)
	assert.Equal(t, 2, s.Len())

	s.Add(3)
	assert.Equal(t, 3, s.Len())
}

func TestSet_All(t *testing.T) {
	s := NewSet[int]()
	assert.Empty(t, slices.Collect(s.All()))

	s.Add(1)
	assert.Equal(t, []int{1}, slices.Collect(s.All()))

	s.Add(2)
	s.Add(2)
	assert.ElementsMatch(t, []int{1, 2}, slices.Collect(s.All()))

	s.Add(3)
	assert.ElementsMatch(t, []int{1, 2, 3}, slices.Collect(s.All()))

	// Stopping early stops the iteration.
	n := 0
	for range s.All() {
		n++
		break
	}
	assert.Equal(t, 1, n)
}

func TestSet_ConstructWithItems(t *testing.T) {
	s := NewSet[int]([]int{1, 2, 2, 3, 4, 4, 6}...)
	assert.Equal(t, 5, s.Len())
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
	assert.Equal(t, 0, s.Len())
	assert.Empty(t, slices.Collect(s.All()))
	assert.True(t, s.Add(1))
	assert.True(t, s.Contains(1))
	assert.Equal(t, 1, s.Len())
}
