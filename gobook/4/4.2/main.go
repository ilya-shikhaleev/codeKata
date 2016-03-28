package main

import "fmt"

func flatten(slice [][]int) []int {
	var result []int
	for _, v := range slice {
		result = append(result, v...)
	}
	return result
}

func main() {
	slice := [][]int{{1, 2, 3, 4, 5}, {6}, {}, {7, 8, 9, 10}}
	noRepeat := flatten(slice)
	fmt.Println(slice)
	fmt.Println(noRepeat)
}
