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

// Map returns a new slice holding f of each element of s, in order.
func Map[T, U any](s []T, f func(T) U) []U {
	res := make([]U, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}

// Reduce folds s into one value, starting from initial and combining in order:
// f(f(f(initial, s[0]), s[1]), ...). With s empty, it returns initial.
func Reduce[T, U any](s []T, f func(U, T) U, initial U) U {
	acc := initial
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// Generate returns a slice of n elements, each a new call of f. Unlike slices.Repeat, which
// copies one value, each element is f's own: a fresh map, pointer or accumulator.
func Generate[T any](f func() T, n int) []T {
	var res = make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = f()
	}
	return res
}

// MergeMaps combines multiple maps into a new map. Values from later maps
// will overwrite values from earlier maps with the same key.
func MergeMaps[K comparable, V any](ms ...map[K]V) map[K]V {
	res := make(map[K]V)
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// Distinct returns the elements of s with each kept once, where it first appears, in a new slice.
// Unlike slices.Compact, duplicates needn't be next to each other.
func Distinct[T comparable](s []T) []T {
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
