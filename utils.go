package gocoll

// Filter returns the elements of slice that f holds for, in order, in a new slice. Unlike
// slices.DeleteFunc, it leaves slice as it was.
func Filter[T any](slice []T, f func(T) bool) []T {
	var res []T
	for _, v := range slice {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// RepeatFunc returns a slice of the given length with elements computed by a function.
func RepeatFunc[T any](f func() T, n int) []T {
	var res = make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = f()
	}
	return res
}

// Merge combines multiple maps into a new map. Values from later maps
// will overwrite values from earlier maps with the same key.
func Merge[K comparable, V any](ms ...map[K]V) map[K]V {
	res := make(map[K]V)
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// Unique returns the elements of s with each kept once, where it first appears, in a new slice.
// Unlike slices.Compact, duplicates needn't be next to each other.
func Unique[T comparable](s []T) []T {
	seen := make(map[T]struct{})
	var res []T
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Partition splits s into the elements pred holds for and the rest, each in s's order.
func Partition[T any](s []T, pred func(T) bool) (pass, fail []T) {
	for _, v := range s {
		if pred(v) {
			pass = append(pass, v)
		} else {
			fail = append(fail, v)
		}
	}
	return pass, fail
}
