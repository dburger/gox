package coll

// MultiSet is a set that counts how many times each value is in it, sometimes called a bag. Its
// zero value is an empty MultiSet, ready to use. A MultiSet shouldn't be copied after first use:
// copies share their counts but not their Len.
type MultiSet[T comparable] struct {
	counts map[T]int
	size   int
}

// NewMultiSet constructs and returns a MultiSet of the given type.
func NewMultiSet[T comparable]() *MultiSet[T] {
	return &MultiSet[T]{}
}

// Add adds an element to the MultiSet and returns its new count.
func (ms *MultiSet[T]) Add(val T) int {
	if ms.counts == nil {
		ms.counts = make(map[T]int)
	}
	ms.counts[val]++
	ms.size++
	return ms.counts[val]
}

// Remove removes one of the value from the MultiSet and returns its new count. Removing a value
// not in the MultiSet changes nothing and returns 0.
func (ms *MultiSet[T]) Remove(val T) int {
	count := ms.counts[val]
	if count == 0 {
		return 0
	}
	ms.size--
	if count == 1 {
		delete(ms.counts, val)
		return 0
	}
	ms.counts[val] = count - 1
	return count - 1
}

// Count returns the count of value in the MultiSet: 0 for a value not in it.
func (ms *MultiSet[T]) Count(val T) int {
	return ms.counts[val]
}

// Len returns the number of elements in the MultiSet, each counted as many times as it was added.
func (ms *MultiSet[T]) Len() int {
	return ms.size
}
