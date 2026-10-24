package coll

import (
	"slices"
	"testing"

	"github.com/dburger/gox/internal/testutil"
)

func makeDeque() *Deque[int] {
	d := NewDeque[int]()
	d.AddFirst(1)
	d.AddLast(2)
	d.AddLast(3)
	d.AddLast(4)
	d.AddLast(5)
	return d
}

func TestDeque_PeekFirst(t *testing.T) {
	d := makeDeque()
	val, ok := d.PeekFirst()
	testutil.Equal(t, ok, true)
	testutil.Equal(t, val, 1)
	val, ok = d.PeekLast()
	testutil.Equal(t, ok, true)
	testutil.Equal(t, val, 5)
}

func TestDeque_PeekLast(t *testing.T) {
	d := makeDeque()
	val, ok := d.PeekLast()
	testutil.Equal(t, ok, true)
	testutil.Equal(t, val, 5)
}

func TestDeque_RemoveFirst(t *testing.T) {
	d := makeDeque()
	val, ok := d.RemoveFirst()
	testutil.Equal(t, ok, true)
	testutil.Equal(t, val, 1)
	testutil.Equal(t, d.Len(), 4)
}

func TestDeque_RemoveLast(t *testing.T) {
	d := makeDeque()
	val, ok := d.RemoveLast()
	testutil.Equal(t, ok, true)
	testutil.Equal(t, val, 5)
	testutil.Equal(t, d.Len(), 4)
}

func TestDeque_IsEmpty(t *testing.T) {
	d := makeDeque()
	testutil.Equal(t, d.IsEmpty(), false)
	d.RemoveFirst()
	d.RemoveLast()
	testutil.Equal(t, d.Len(), 3)
	d.RemoveFirst()
	d.RemoveLast()
	d.RemoveFirst()
	testutil.Equal(t, d.IsEmpty(), true)
}

func TestDeque_ZeroValue(t *testing.T) {
	var d Deque[int]
	testutil.Equal(t, d.IsEmpty(), true)
	testutil.Equal(t, d.Len(), 0)
	for _, f := range []func() (int, bool){d.PeekFirst, d.PeekLast, d.RemoveFirst, d.RemoveLast} {
		v, ok := f()
		testutil.Equal(t, ok, false)
		testutil.Equal(t, v, 0)
	}
	for range d.All() {
		t.Error("iterated over an element of an empty Deque")
	}

	d.AddLast(2)
	d.AddFirst(1)
	testutil.EqualSlices(t, slices.Collect(d.All()), []int{1, 2})
	v, ok := d.RemoveLast()
	testutil.Equal(t, ok, true)
	testutil.Equal(t, v, 2)
}

func TestDeque_All(t *testing.T) {
	d := makeDeque()
	testutil.EqualSlices(t, slices.Collect(d.All()), []int{1, 2, 3, 4, 5})
	var first []int
	for v := range d.All() {
		first = append(first, v)
		if len(first) == 2 {
			break
		}
	}
	testutil.EqualSlices(t, first, []int{1, 2})
}
