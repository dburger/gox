package slicesx

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6}
	filtered := Filter(input, func(i int) bool { return i%2 == 0 })
	expected := []int{0, 2, 4, 6}
	assert.Equal(t, expected, filtered)
}

func TestGenerate(t *testing.T) {
	expected := []int{42, 42, 42, 42, 42, 42}
	fortyTwos := Generate(func() int { return 42 }, 6)

	assert.Equal(t, expected, fortyTwos)
}

func TestDistinct(t *testing.T) {
	slice := []int{2, 1, 2, 2, 3, 1, 4, 4, 5}
	expected := []int{2, 1, 3, 4, 5}
	assert.Equal(t, expected, Distinct(slice))
}

func TestPartition(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	pass, fail := Partition(slice, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 4}, pass)
	assert.Equal(t, []int{1, 3, 5}, fail)
}

func TestPartition_WhenEmpty(t *testing.T) {
	slice := []int{}
	pass, fail := Partition(slice, func(i int) bool { return i%2 == 0 })
	assert.Nil(t, pass)
	assert.Nil(t, fail)
}

func TestMap(t *testing.T) {
	assert.Equal(t, []int{2, 4, 6, 8}, Map([]int{1, 2, 3, 4}, func(i int) int { return i * 2 }))
	assert.Equal(t, []string{"1", "22"}, Map([]int{1, 22}, strconv.Itoa))
	assert.Empty(t, Map([]int{}, strconv.Itoa))
	assert.Empty(t, Map(nil, strconv.Itoa))
}

func TestReduce(t *testing.T) {
	items := []int{1, 2, 3}
	assert.Equal(t, 6, Reduce(items, func(a, b int) int { return a + b }, 0))
	assert.Equal(t, 6, Reduce(items, func(a, b int) int { return a * b }, 1))
	// The order is left to right, and the result's type can differ from the elements'.
	assert.Equal(t, "x123", Reduce(items, func(acc string, i int) string { return acc + strconv.Itoa(i) }, "x"))
	assert.Equal(t, 7, Reduce(nil, func(a, b int) int { return a + b }, 7))
}
