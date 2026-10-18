package gocoll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var SAMPLE_MAP = map[string]string{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
}

func TestFilter(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6}
	filtered := Filter(input, func(i int) bool { return i%2 == 0 })
	expected := []int{0, 2, 4, 6}
	assert.Equal(t, expected, filtered)
}

func TestRepeatFunc(t *testing.T) {
	expected := []int{42, 42, 42, 42, 42, 42}
	fortyTwos := RepeatFunc(func() int { return 42 }, 6)

	assert.Equal(t, expected, fortyTwos)
}

func TestRepeatValues(t *testing.T) {
	expected := []int{6, 6, 6}
	sixes := RepeatValues(6, 3)

	assert.Equal(t, expected, sixes)
}

func TestMapKeys(t *testing.T) {
	keys := MapKeys(SAMPLE_MAP)
	assert.Len(t, keys, len(keys))
	for _, key := range []string{"key1", "key2", "key3"} {
		assert.Contains(t, keys, key)
	}
}

func TestMapValues(t *testing.T) {
	values := MapValues(SAMPLE_MAP)
	assert.Len(t, values, len(values))
	for _, value := range []string{"value1", "value2", "value3"} {
		assert.Contains(t, values, value)
	}
}

func TestMerge(t *testing.T) {
	map1 := map[string]string{"key1": "value1", "key2": "value2"}
	map2 := map[string]string{"key2": "value3", "key4": "value4"}
	map3 := map[string]string{"key5": "value5"}
	merged := Merge(map1, map2, map3)
	assert.Equal(t, 4, len(merged))
	assert.Equal(t, "value1", merged["key1"])
	assert.Equal(t, "value3", merged["key2"])
	assert.Equal(t, "value4", merged["key4"])
	assert.Equal(t, "value5", merged["key5"])
	// TODO(dburger): could assert original maps unhanged
}

func TestMergeInto(t *testing.T) {
	map1 := map[string]string{"key1": "value1", "key2": "value2"}
	map2 := map[string]string{"key2": "value3", "key4": "value4"}
	map3 := map[string]string{"key5": "value5"}
	MergeInto(map1, map2, map3)
	assert.Equal(t, 4, len(map1))
	assert.Equal(t, "value1", map1["key1"])
	assert.Equal(t, "value3", map1["key2"])
	assert.Equal(t, "value4", map1["key4"])
	assert.Equal(t, "value5", map1["key5"])
	// TODO(dburger): could assert original maps unhanged
}

func TestUnique(t *testing.T) {
	slice := []int{2, 1, 2, 2, 3, 1, 4, 4, 5}
	expected := []int{2, 1, 3, 4, 5}
	assert.Equal(t, expected, Unique(slice))
}
