package sort

import "sort"

// func BubbleSort(elements []int) {
// 	keepWorking := true
// 	for keepWorking {
// 		keepWorking = false
// 		for i := 0; i < len(elements)-1; i++ {
// 			if elements[i] > elements[i+1] {
// 				keepWorking = true
// 				elements[i], elements[i+1] = elements[i+1], elements[i]
// 			}
// 		}
// 	}
// }

func Sort(elements []int) {
	sort.Ints(elements)
}

func BubbleSort(arr []int) []int {
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				flag = true
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

func GenerateArr(num int) []int {
	n := make([]int, num)
	j := 0
	for i := num - 1; i >= 0; i-- {
		n[j] = i
		j++
	}
	return n
}
