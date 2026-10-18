package coll

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	val, ok = d.PeekLast()
	assert.True(t, ok)
	assert.Equal(t, 5, val)
}

func TestDeque_PeekLast(t *testing.T) {
	d := makeDeque()
	val, ok := d.PeekLast()
	assert.True(t, ok)
	assert.Equal(t, 5, val)
}

func TestDeque_RemoveFirst(t *testing.T) {
	d := makeDeque()
	val, ok := d.RemoveFirst()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 4, d.Len())
}

func TestDeque_RemoveLast(t *testing.T) {
	d := makeDeque()
	val, ok := d.RemoveLast()
	assert.True(t, ok)
	assert.Equal(t, 5, val)
	assert.Equal(t, 4, d.Len())
}

func TestDeque_IsEmpty(t *testing.T) {
	d := makeDeque()
	assert.False(t, d.IsEmpty())
	d.RemoveFirst()
	d.RemoveLast()
	assert.Equal(t, 3, d.Len())
	d.RemoveFirst()
	d.RemoveLast()
	d.RemoveFirst()
	assert.True(t, d.IsEmpty())
}

func TestDeque_ZeroValue(t *testing.T) {
	var d Deque[int]
	assert.True(t, d.IsEmpty())
	assert.Equal(t, 0, d.Len())
	for _, f := range []func() (int, bool){d.PeekFirst, d.PeekLast, d.RemoveFirst, d.RemoveLast} {
		v, ok := f()
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	}
	for range d.All() {
		t.Error("iterated over an element of an empty Deque")
	}

	d.AddLast(2)
	d.AddFirst(1)
	assert.Equal(t, []int{1, 2}, slices.Collect(d.All()))
	v, ok := d.RemoveLast()
	assert.True(t, ok)
	assert.Equal(t, 2, v)
}

func TestDeque_All(t *testing.T) {
	d := makeDeque()
	assert.Equal(t, []int{1, 2, 3, 4, 5}, slices.Collect(d.All()))
	var first []int
	for v := range d.All() {
		first = append(first, v)
		if len(first) == 2 {
			break
		}
	}
	assert.Equal(t, []int{1, 2}, first)
}
