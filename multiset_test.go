package gocoll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiSet_Add(t *testing.T) {
	ms := NewMultiSet[int]()
	count := ms.Add(1)
	assert.Equal(t, 1, count)
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 0, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	count = ms.Add(2)
	assert.Equal(t, 1, count)
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 1, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	count = ms.Add(2)
	assert.Equal(t, 2, count)
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 2, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	count = ms.Add(3)
	assert.Equal(t, 1, count)
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 2, ms.Count(2))
	assert.Equal(t, 1, ms.Count(3))
	assert.Equal(t, 0, ms.Count(4))
}

func TestMultiSet_Remove(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Add(1)
	ms.Add(2)
	ms.Add(2)
	ms.Add(3)

	assert.Equal(t, 4, ms.Size())

	count := ms.Remove(4)
	assert.Equal(t, 0, count)
	assert.Equal(t, 4, ms.Size())

	count = ms.Remove(3)
	assert.Equal(t, 0, count)
	assert.Equal(t, 3, ms.Size())

	count = ms.Remove(2)
	assert.Equal(t, 1, count)
	assert.Equal(t, 2, ms.Size())

	count = ms.Remove(2)
	assert.Equal(t, 0, count)
	assert.Equal(t, 1, ms.Size())

	count = ms.Remove(2)
	assert.Equal(t, 0, count)
	assert.Equal(t, 1, ms.Size())

	count = ms.Remove(1)
	assert.Equal(t, 0, count)
	assert.Equal(t, 0, ms.Size())
}

func TestMultiSet_Count(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Add(1)
	ms.Add(2)
	ms.Add(2)
	ms.Add(3)

	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 2, ms.Count(2))
	assert.Equal(t, 1, ms.Count(3))
	assert.Equal(t, 0, ms.Count(4))
}

func TestMultiSet_Size(t *testing.T) {
	ms := NewMultiSet[int]()
	assert.Equal(t, 0, ms.Size())

	ms.Add(1)
	assert.Equal(t, 1, ms.Size())

	ms.Add(2)
	assert.Equal(t, 2, ms.Size())

	ms.Add(2)
	assert.Equal(t, 3, ms.Size())

	ms.Add(3)
	assert.Equal(t, 4, ms.Size())
}

func TestMultiSet_ZeroValue(t *testing.T) {
	var ms MultiSet[int]
	assert.Equal(t, 0, ms.Count(1))
	assert.Equal(t, 0, ms.Remove(1))
	assert.Equal(t, 0, ms.Size())
	assert.Equal(t, 1, ms.Add(1))
	assert.Equal(t, 2, ms.Add(1))
	assert.Equal(t, 2, ms.Size())
}
