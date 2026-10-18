package gocoll

import (
	"maps"
	"slices"
)

// Set provides a set type. Its zero value is an empty set, ready to use.
type Set[T comparable] struct {
	// Idiomatic approach, struct{} does not require an allocation
	// and communicates no meaning.
	items map[T]struct{}
}

// NewSet constructs and returns a set of the given type, holding items.
func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{}
	s.AddAll(items...)
	return s
}

// Contains returns whether the Set contains the indicated item.
func (s *Set[T]) Contains(i T) bool {
	_, ok := s.items[i]
	return ok
}

// Add adds the given item to the Set and returns whether the item was added.
// If the item was already present, false is returned.
func (s *Set[T]) Add(i T) bool {
	if s.items == nil {
		s.items = make(map[T]struct{})
	}
	_, ok := s.items[i]
	s.items[i] = struct{}{}
	return !ok
}

// AddAll adds each item of items to the Set.
func (s *Set[T]) AddAll(items ...T) {
	for _, i := range items {
		s.Add(i)
	}
}

// Remove removes the given item from the Set and returns whether the item was removed.
// If it was not present, false is returned.
func (s *Set[T]) Remove(i T) bool {
	_, ok := s.items[i]
	if ok {
		delete(s.items, i)
	}
	return ok
}

// Size returns a count of the items in the Set.
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Slice returns the contents of the set as a slice.
func (s *Set[T]) Slice() []T {
	return slices.Collect(maps.Keys(s.items))
}
