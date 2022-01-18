package lc15

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	rst := threeSumTwoPointer(input)

	exp := [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}
	if !arrEqual(exp, rst) {
		t.Errorf("expected %v is different from actual %v", exp, rst)
	}

}

func arrEqual(ar1, ar2 [][]int) bool {
	return fmt.Sprint(ar1) == fmt.Sprint(ar2)
}
