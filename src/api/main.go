package main

import (
	"fmt"

	"github.com/stepup99/golang-testing/src/api/utils/sort"
)

func init() {
	// gin.SetMode(gin.ReleaseMode)
}
func main() {
	// country, err := locations_provider.GetCountry("AR")
	// fmt.Println(err)
	// fmt.Println(country)
	// this is simple server getting created
	// app.StartApp()

	// arr := sort.BubbleSort([]int{1, 23, 32, 5})
	arr := sort.GenerateArr(10)
	fmt.Println(arr)
}
