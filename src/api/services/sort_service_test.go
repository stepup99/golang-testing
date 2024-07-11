package services

import (
	"testing"
)

func TestSortService(t *testing.T) {
	elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}
	// Execution
	Sort(elements)
	// Validation

	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last elements should be 0")
	}
}
