package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tcs := []struct {
		sls []int
		tgt int
		exp []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range tcs {
		rst := twoSum(tc.sls, tc.tgt)
		if !Equal(tc.exp, rst) {
			t.Errorf("expected %v but got %v", tc.exp, rst)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	tcs := []struct {
		sls []int
		tgt int
		exp []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range tcs {
		rst := twoSum2(tc.sls, tc.tgt)
		if !Equal(tc.exp, rst) {
			t.Errorf("expected %v but got %v", tc.exp, rst)
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
