package binchop1

// Returns the integer index of the target in the array, or -1 if the target is not in the array
type interval struct {
	arr    []int
	top    int
	bottom int
}

func getMiddle(top int, bottom int) int {
	return int((top + bottom) / 2)
}

func newInterval(arr []int) *interval {
	return &interval{arr, len(arr) - 1, 0}
}

func Chop(x int, arr []int) int {
	result := -1
	if len(arr) == 0 { //initially I forgot about this.
		return result
	}

	inter := newInterval(arr)
	for {
		if inter.top == inter.bottom {
			break
		}
		middle := getMiddle(inter.top, inter.bottom)
		if inter.arr[middle] < x {
			if inter.bottom == middle {
				inter.bottom = inter.top
				break
			}
			inter.bottom = middle
		} else {
			inter.top = middle
		}
	}

	if inter.arr[inter.top] == x {
		result = inter.top
	}

	return result
}
