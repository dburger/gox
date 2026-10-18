package gocoll

// MultiSet provides a map that keeps counts of elements. Its zero value is empty, ready to use.
type MultiSet[T comparable] struct {
	counts map[T]int
	size   int
}

// NewMultiSet constructs and returns a MultiSet of the given type.
func NewMultiSet[T comparable]() *MultiSet[T] {
	return &MultiSet[T]{make(map[T]int), 0}
}

// Add adds an element to the multiset and returns the new count.
func (ms *MultiSet[T]) Add(val T) int {
	if ms.counts == nil {
		ms.counts = make(map[T]int)
	}
	count, ok := ms.counts[val]
	if !ok {
		count = 1
	} else {
		count++
	}
	ms.counts[val] = count
	ms.size += 1
	return count
}

// Remove removes the value from the MultiSet and returns the new count.
func (ms *MultiSet[T]) Remove(val T) int {
	count, ok := ms.counts[val]
	if !ok {
		// Nothing to do, not there.
	} else {
		ms.size--
		count--
		if count == 0 {
			delete(ms.counts, val)
		} else {
			ms.counts[val] = count
		}
	}
	return count
}

// Count returns the count of value in the MultiSet.
func (ms *MultiSet[T]) Count(val T) int {
	count, ok := ms.counts[val]
	if !ok {
		count = 0
	}
	return count
}

func (ms *MultiSet[T]) Size() int {
	return ms.size
}
