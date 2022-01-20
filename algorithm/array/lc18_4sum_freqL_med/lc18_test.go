package lc18

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	tcs := []struct {
		arr    []int
		target int
		exp    [][]int
	}{
		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
				[]int{-1, 0, 0, 1},
			},
		},
		{
			[]int{2, 2, 2, 2, 2},
			8,
			[][]int{
				[]int{2, 2, 2, 2},
			},
		},
	}

	for _, tc := range tcs {
		rst := fourSum(tc.arr, tc.target)
		if fmt.Sprint(rst) != fmt.Sprint(tc.exp) {
			t.Errorf("expecting %v but got %v", tc.exp, rst)
		}

	}
}
