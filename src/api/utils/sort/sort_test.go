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

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last elements should be 9")
	}
	fmt.Println(elements)
	// run this command PS C:\Users\HP\Documents\GitHub\go_projects\golang-testing> go test -v ./...
	// PS C:\Users\HP\Documents\GitHub\go_projects\golang-testing> go test -cover ./...
	// PS C:\Users\HP\Documents\GitHub\go_projects\golang-testing> go test ./utils/sort -coverprofile coverage.out
	// PS C:\Users\HP\Documents\GitHub\go_projects\golang-testing> go tool cover -html coverage.out
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := getElemnts(10000)
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	// Execution
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
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

func BenchmarkSort(b *testing.B) {
	elements := getElemnts(10000)
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

// benchmarksort
// 64249             19009 ns/op               1 B/op          0 allocs/op
// 59576             19625 ns/op               1 B/op          0 allocs/op

// benchmarkbubblesort

// 171352              5923 ns/op               0 B/op          0 allocs/op
// 219704              5649 ns/op               0 B/op          0 allocs/op


// PS C:\Users\HP\Documents\GitHub\go_projects\golang-testing\src\api\utils\sort> go test -bench BenchmarkBubbleSort