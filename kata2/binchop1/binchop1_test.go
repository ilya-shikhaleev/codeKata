package binchop1

import "testing"

func TestGetMiddle(t *testing.T) {
	cases := []struct {
		top    int
		bottom int
		want   int
	}{
		{0, 0, 0},
		{1, 0, 0},
		{2, 0, 1},
		{3, 0, 1},
		{4, 0, 2},
		{5, 0, 2},
		{6, 0, 3},

		{10, 5, 7},
		{10, 2, 6},
		{10, 8, 9},
		{10, 9, 9},
	}

	for _, c := range cases {
		middle := getMiddle(c.top, c.bottom)
		if middle != c.want {
			t.Errorf("getMiddle(%v, %v) = %v, want %v", c.top, c.bottom, middle, c.want)
		}
	}
}

func TestNewInterval(t *testing.T) {
	inter := newInterval([]int{1, 2, 3, 4})
	if inter.top != 3 {
		t.Errorf("newInterval([]int{1,2,3,4}) creates %v, want top = 3")
	}
	if inter.bottom != 0 {
		t.Errorf("newInterval([]int{1,2,3,4}) creates %v, want bottom = 0")
	}
}
