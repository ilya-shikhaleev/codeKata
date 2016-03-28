package main

import "fmt"

func removeRepeats(slice []int) []int {
	result := make([]int, 0, len(slice))
	uniq := make(map[int]bool, len(slice))
	for _, v := range slice {
		if _, ok := uniq[v]; !ok {
			result = append(result, v)
			uniq[v] = true
		}
	}
	return result
}

func main() {
	slice := []int{9, 9, 1, 9, 2, 3, 4, 5, 1, 34, 4, 5, 3, 5, 26, 5, 6, 594, 62, 45, 62, 3, 4, 5, 345}
	noRepeat := removeRepeats(slice)
	fmt.Println(slice)
	fmt.Println(noRepeat)
}
