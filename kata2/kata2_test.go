package main

import (
	"testing"
)

func TestChop(t *testing.T) {
	cases := []struct {
		x    int
		arr  []int
		want int
	}{
		{3, []int{}, -1},
		{3, []int{1}, -1},
		{1, []int{1}, 0},

		{1, []int{1, 3, 5}, 0},
		{3, []int{1, 3, 5}, 1},
		{5, []int{1, 3, 5}, 2},
		{0, []int{1, 3, 5}, -1},
		{2, []int{1, 3, 5}, -1},
		{4, []int{1, 3, 5}, -1},
		{6, []int{1, 3, 5}, -1},

		{1, []int{1, 3, 5, 7}, 0},
		{3, []int{1, 3, 5, 7}, 1},
		{5, []int{1, 3, 5, 7}, 2},
		{7, []int{1, 3, 5, 7}, 3},
		{0, []int{1, 3, 5, 7}, -1},
		{2, []int{1, 3, 5, 7}, -1},
		{4, []int{1, 3, 5, 7}, -1},
		{6, []int{1, 3, 5, 7}, -1},
		{8, []int{1, 3, 5, 7}, -1},
	}
	for _, c := range cases {
		got := Chop(c.x, c.arr)
		if got != c.want {
			t.Errorf("Chop(%v, %v) == %v, want %v", c.x, c.arr, got, c.want)
		}
	}
}
