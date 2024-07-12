package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	// init
	// tests := []struct {
	// 	input    []int
	// 	expected []int
	// }{
	// 	{[]int{3, 2, 1}, []int{1, 2, 3}},
	// 	{[]int{}, []int{}},
	// }
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	// for _, test := range tests {
	// elements := getElemnts(10)
	// Execution

	// if len(elemnts) != 10 {
	// 	t.Error("length should be 10")
	// }
	assert.Nil(t, elements, "length should be there ")
	assert.EqualValues(t, 1, len(elements), "length should be 10")

	assert.EqualValues(t, 7, elements[1], "first element should be 7")

	assert.EqualValues(t, 91, elements[0], "first element should be 9")
	BubbleSort(elements)
	// Validation

	// if len(elemnts) != 9 {
	// 	t.Error("length should be 9")
	// }

	// if elements[len(elements)-1] != 9 {
	// 	t.Error("last elements should be 9")
	// }

	// }

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
