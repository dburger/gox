package coll

import (
	"slices"
	"testing"

	"github.com/dburger/gox/internal/testutil"
)

func TestSet_Contains(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	testutil.Equal(t, s.Contains(1), true)
	testutil.Equal(t, s.Contains(2), true)
	testutil.Equal(t, s.Contains(3), true)
	testutil.Equal(t, s.Contains(4), false)
}

func TestSet_Add(t *testing.T) {
	s := NewSet[int]()
	added := s.Add(1)
	testutil.Equal(t, added, true)
	testutil.Equal(t, s.Len(), 1)
	testutil.Equal(t, s.Contains(1), true)
	testutil.Equal(t, s.Contains(2), false)
	testutil.Equal(t, s.Contains(3), false)

	added = s.Add(2)
	testutil.Equal(t, added, true)
	testutil.Equal(t, s.Len(), 2)
	testutil.Equal(t, s.Contains(1), true)
	testutil.Equal(t, s.Contains(2), true)
	testutil.Equal(t, s.Contains(3), false)

	added = s.Add(2)
	testutil.Equal(t, added, false)
	testutil.Equal(t, s.Len(), 2)
	testutil.Equal(t, s.Contains(1), true)
	testutil.Equal(t, s.Contains(2), true)
	testutil.Equal(t, s.Contains(3), false)

	added = s.Add(3)
	testutil.Equal(t, added, true)
	testutil.Equal(t, s.Len(), 3)
	testutil.Equal(t, s.Contains(1), true)
	testutil.Equal(t, s.Contains(2), true)
	testutil.Equal(t, s.Contains(3), true)
}

func TestSet_AddAll(t *testing.T) {
	s := NewSet[int]()
	s.AddAll([]int{1, 2, 2, 3, 4, 4, 6}...)
	testutil.Equal(t, s.Len(), 5)
	testutil.Equal(t, s.Contains(1), true)
	testutil.Equal(t, s.Contains(2), true)
	testutil.Equal(t, s.Contains(3), true)
	testutil.Equal(t, s.Contains(4), true)
	testutil.Equal(t, s.Contains(5), false)
	testutil.Equal(t, s.Contains(6), true)
}

func TestSet_Remove(t *testing.T) {
	s := NewSet[int]()
	testutil.Equal(t, s.Len(), 0)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	testutil.Equal(t, s.Len(), 3)

	s.Remove(1)
	testutil.Equal(t, s.Len(), 2)
	testutil.Equal(t, s.Contains(1), false)
	testutil.Equal(t, s.Contains(2), true)
	testutil.Equal(t, s.Contains(3), true)

	s.Remove(2)
	testutil.Equal(t, s.Len(), 1)
	testutil.Equal(t, s.Contains(1), false)
	testutil.Equal(t, s.Contains(2), false)
	testutil.Equal(t, s.Contains(3), true)

	s.Remove(3)
	testutil.Equal(t, s.Len(), 0)
	testutil.Equal(t, s.Contains(1), false)
	testutil.Equal(t, s.Contains(2), false)
	testutil.Equal(t, s.Contains(3), false)
}

func TestSet_Size(t *testing.T) {
	s := NewSet[int]()
	testutil.Equal(t, s.Len(), 0)

	s.Add(1)
	testutil.Equal(t, s.Len(), 1)

	s.Add(2)
	testutil.Equal(t, s.Len(), 2)

	s.Add(2)
	testutil.Equal(t, s.Len(), 2)

	s.Add(3)
	testutil.Equal(t, s.Len(), 3)
}

func TestSet_All(t *testing.T) {
	s := NewSet[int]()
	testutil.Equal(t, len(slices.Collect(s.All())), 0)

	s.Add(1)
	testutil.EqualSlices(t, slices.Collect(s.All()), []int{1})

	s.Add(2)
	s.Add(2)
	testutil.EqualSlices(t, slices.Sorted(s.All()), []int{1, 2})

	s.Add(3)
	testutil.EqualSlices(t, slices.Sorted(s.All()), []int{1, 2, 3})

	// Stopping early stops the iteration.
	n := 0
	for range s.All() {
		n++
		break
	}
	testutil.Equal(t, n, 1)
}

func TestSet_ConstructWithItems(t *testing.T) {
	s := NewSet[int]([]int{1, 2, 2, 3, 4, 4, 6}...)
	testutil.Equal(t, s.Len(), 5)
	testutil.Equal(t, s.Contains(1), true)
	testutil.Equal(t, s.Contains(2), true)
	testutil.Equal(t, s.Contains(3), true)
	testutil.Equal(t, s.Contains(4), true)
	testutil.Equal(t, s.Contains(6), true)
}

func TestSet_ZeroValue(t *testing.T) {
	var s Set[int]
	testutil.Equal(t, s.Contains(1), false)
	testutil.Equal(t, s.Remove(1), false)
	testutil.Equal(t, s.Len(), 0)
	testutil.Equal(t, len(slices.Collect(s.All())), 0)
	testutil.Equal(t, s.Add(1), true)
	testutil.Equal(t, s.Contains(1), true)
	testutil.Equal(t, s.Len(), 1)
}
