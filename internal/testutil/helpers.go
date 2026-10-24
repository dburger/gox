// Package testutil holds helpers for gox's tests.
package testutil

import (
	"slices"
	"testing"
)

// Equal reports an error if got isn't want.
func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

// EqualSlices reports an error if got doesn't hold want's elements in want's order.
func EqualSlices[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
