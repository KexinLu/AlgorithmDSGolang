package lc16

/**
https://leetcode.com/problems/3sum-closest/
**/

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	best := math.MaxInt64
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				return sum
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(best-target)) {
				best = sum
			}
		}
	}
	return best

}
