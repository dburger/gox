package coll

import (
	"testing"

	"github.com/dburger/gox/internal/testutil"
)

func TestMultiSet_Add(t *testing.T) {
	ms := NewMultiSet[int]()
	count := ms.Add(1)
	testutil.Equal(t, count, 1)
	testutil.Equal(t, ms.Count(1), 1)
	testutil.Equal(t, ms.Count(2), 0)
	testutil.Equal(t, ms.Count(3), 0)
	count = ms.Add(2)
	testutil.Equal(t, count, 1)
	testutil.Equal(t, ms.Count(1), 1)
	testutil.Equal(t, ms.Count(2), 1)
	testutil.Equal(t, ms.Count(3), 0)
	count = ms.Add(2)
	testutil.Equal(t, count, 2)
	testutil.Equal(t, ms.Count(1), 1)
	testutil.Equal(t, ms.Count(2), 2)
	testutil.Equal(t, ms.Count(3), 0)
	count = ms.Add(3)
	testutil.Equal(t, count, 1)
	testutil.Equal(t, ms.Count(1), 1)
	testutil.Equal(t, ms.Count(2), 2)
	testutil.Equal(t, ms.Count(3), 1)
	testutil.Equal(t, ms.Count(4), 0)
}

func TestMultiSet_Remove(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Add(1)
	ms.Add(2)
	ms.Add(2)
	ms.Add(3)

	testutil.Equal(t, ms.Len(), 4)

	count := ms.Remove(4)
	testutil.Equal(t, count, 0)
	testutil.Equal(t, ms.Len(), 4)

	count = ms.Remove(3)
	testutil.Equal(t, count, 0)
	testutil.Equal(t, ms.Len(), 3)

	count = ms.Remove(2)
	testutil.Equal(t, count, 1)
	testutil.Equal(t, ms.Len(), 2)

	count = ms.Remove(2)
	testutil.Equal(t, count, 0)
	testutil.Equal(t, ms.Len(), 1)

	count = ms.Remove(2)
	testutil.Equal(t, count, 0)
	testutil.Equal(t, ms.Len(), 1)

	count = ms.Remove(1)
	testutil.Equal(t, count, 0)
	testutil.Equal(t, ms.Len(), 0)
}

func TestMultiSet_Count(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Add(1)
	ms.Add(2)
	ms.Add(2)
	ms.Add(3)

	testutil.Equal(t, ms.Count(1), 1)
	testutil.Equal(t, ms.Count(2), 2)
	testutil.Equal(t, ms.Count(3), 1)
	testutil.Equal(t, ms.Count(4), 0)
}

func TestMultiSet_Size(t *testing.T) {
	ms := NewMultiSet[int]()
	testutil.Equal(t, ms.Len(), 0)

	ms.Add(1)
	testutil.Equal(t, ms.Len(), 1)

	ms.Add(2)
	testutil.Equal(t, ms.Len(), 2)

	ms.Add(2)
	testutil.Equal(t, ms.Len(), 3)

	ms.Add(3)
	testutil.Equal(t, ms.Len(), 4)
}

func TestMultiSet_ZeroValue(t *testing.T) {
	var ms MultiSet[int]
	testutil.Equal(t, ms.Count(1), 0)
	testutil.Equal(t, ms.Remove(1), 0)
	testutil.Equal(t, ms.Len(), 0)
	testutil.Equal(t, ms.Add(1), 1)
	testutil.Equal(t, ms.Add(1), 2)
	testutil.Equal(t, ms.Len(), 2)
}
