package main

import (
	"fmt"
	"math"
)

func make2D(slice []int, size int) [][]int {
	resultSize := len(slice) / size
	if len(slice)%size > 0 {
		resultSize++
	}
	result := make([][]int, resultSize)
	for i := 0; i < resultSize; i++ {
		result[i] = make([]int, size)
		rightBound := int(math.Min(float64((i+1)*size), float64(len(slice))))
		copy(result[i], slice[i*size:rightBound])
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(slice)
	fmt.Println(make2D(slice, 3))
	fmt.Println(make2D(slice, 4))
	fmt.Println(make2D(slice, 5))
	fmt.Println(make2D(slice, 6))
}
