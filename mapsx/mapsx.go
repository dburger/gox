// Package mapsx provides map functions the standard library's maps package doesn't.
package mapsx

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
