package services

import (
	"testing"
)

func TestSortService(t *testing.T) {
	elements := getElemnts(10)
	// elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}
	// Execution
	Sort(elements)
	// Validation

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last elements should be 9")
	}
}

func TestSortServiceMoreThan10000(t *testing.T) {
	elements := getElemnts(10001)
	// elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}
	// Execution
	Sort(elements)
	// Validation

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 10000 {
		t.Error("last elements should be 10000")
	}
}

func getElemnts(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}

// PS C:\Users\HP\Documents\GitHub\go_projects\golang-testing> go test -run TestSortService ./src/api/services -v
