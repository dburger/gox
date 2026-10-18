package gocoll

import (
	"iter"
	"maps"
)

// Set is a set of comparable values, each held at most once. Its zero value is an empty Set,
// ready to use. A Set shouldn't be copied after first use: copies share their items.
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

// Len returns the number of items in the Set.
func (s *Set[T]) Len() int {
	return len(s.items)
}

// All returns an iterator over the Set's items, in no particular order: for item := range s.All().
// slices.Collect(s.All()) gives them as a slice.
func (s *Set[T]) All() iter.Seq[T] {
	return maps.Keys(s.items)
}
