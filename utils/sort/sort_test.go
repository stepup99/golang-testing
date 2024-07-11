package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	// Execution
	BubbleSort(elements)
	// Validation

	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last elements should be 0")
	}
	fmt.Println(elements)
}
