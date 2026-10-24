package slicesx

import (
	"strconv"
	"testing"

	"github.com/dburger/gox/internal/testutil"
)

func TestFilter(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6}
	filtered := Filter(input, func(i int) bool { return i%2 == 0 })
	expected := []int{0, 2, 4, 6}
	testutil.EqualSlices(t, filtered, expected)
}

func TestGenerate(t *testing.T) {
	expected := []int{42, 42, 42, 42, 42, 42}
	fortyTwos := Generate(func() int { return 42 }, 6)

	testutil.EqualSlices(t, fortyTwos, expected)
}

func TestDistinct(t *testing.T) {
	slice := []int{2, 1, 2, 2, 3, 1, 4, 4, 5}
	expected := []int{2, 1, 3, 4, 5}
	testutil.EqualSlices(t, Distinct(slice), expected)
}

func TestPartition(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	pass, fail := Partition(slice, func(i int) bool { return i%2 == 0 })
	testutil.EqualSlices(t, pass, []int{2, 4})
	testutil.EqualSlices(t, fail, []int{1, 3, 5})
}

func TestPartition_WhenEmpty(t *testing.T) {
	slice := []int{}
	pass, fail := Partition(slice, func(i int) bool { return i%2 == 0 })
	if pass != nil || fail != nil {
		t.Errorf("got %v and %v, want nil and nil", pass, fail)
	}
}

func TestMap(t *testing.T) {
	testutil.EqualSlices(t, Map([]int{1, 2, 3, 4}, func(i int) int { return i * 2 }), []int{2, 4, 6, 8})
	testutil.EqualSlices(t, Map([]int{1, 22}, strconv.Itoa), []string{"1", "22"})
	testutil.Equal(t, len(Map([]int{}, strconv.Itoa)), 0)
	testutil.Equal(t, len(Map(nil, strconv.Itoa)), 0)
}

func TestReduce(t *testing.T) {
	items := []int{1, 2, 3}
	testutil.Equal(t, Reduce(items, func(a, b int) int { return a + b }, 0), 6)
	testutil.Equal(t, Reduce(items, func(a, b int) int { return a * b }, 1), 6)
	// The order is left to right, and the result's type can differ from the elements'.
	testutil.Equal(t, Reduce(items, func(acc string, i int) string { return acc + strconv.Itoa(i) }, "x"), "x123")
	testutil.Equal(t, Reduce(nil, func(a, b int) int { return a + b }, 7), 7)
}
