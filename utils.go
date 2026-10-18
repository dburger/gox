package gocoll

// Supposedly slices.Repeat coming in 1.23.

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

// RepeatValues returns a slice of the given length with elements of the given value.
func RepeatValues[T any](val T, n int) []T {
	var res = make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = val
	}
	return res
}

// MapKeys returns a slice of the keys in the map.
// Alternatively available from golang.org/x/exp/maps.
func MapKeys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValues returns a slice of the values in the map.
// Alternatively available from golang.org/x/exp/maps.
func MapValues[T comparable, U any](m map[T]U) []U {
	values := make([]U, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Merge combines multiple maps into a new map. Values from later maps
// will overwrite values from earlier maps with the same key.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	res := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// MergeInto merges multiple maps into the destination map. Values from later maps
// will overwrite values from earlier maps with the same key.
func MergeInto[K comparable, V any](dest map[K]V, maps ...map[K]V) {
	for _, m := range maps {
		for k, v := range m {
			dest[k] = v
		}
	}
}

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
