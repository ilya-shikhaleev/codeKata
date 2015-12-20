package main

import (
	"fmt"
	"github.com/ilya-shikhaleev/codeKata/kata2/binchop1"
)

// Returns the integer index of the target in the array, or -1 if the target is not in the array
func Chop(x int, arr []int) int {
	return binchop1.Chop(x, arr)
}

func main() {
	x := 10
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	index := Chop(x, arr)
	fmt.Println(index)
}
