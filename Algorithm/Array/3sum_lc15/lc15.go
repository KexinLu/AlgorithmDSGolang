package lc15

import "sort"

/**
Given an integer array nums, return **all** the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must **not contain duplicate** triplets.
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Example 2:

Input: nums = []
Output: []
Example 3:

Input: nums = [0]
Output: []

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

https://leetcode.com/problems/3sum/
**/

// Solution with marker map
func threeSumSolution1(nums []int) [][]int {
	output := [][]int{}
	outputMap := make(map[[3]int][]int)
	n := len(nums)
	if n < 3 {
		return output
	}

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		marker := make(map[int]int)
		for j := i + 1; j < n; j++ {
			curTarget := target - nums[j]
			if val, ok := marker[nums[j]]; ok {
				outputMap[[3]int{nums[i], nums[j], val}] = []int{nums[i], nums[j], val}
			} else {
				marker[curTarget] = nums[j]
			}
		}
	}
	for _, v := range outputMap {
		output = append(output, v)
	}

	return output
}

// better solution in this case
func threeSumTwoPointer(nums []int) [][]int {
	output, n := [][]int{}, len(nums)

	// sourt the input so we can skip same item and use two pointers
	sort.Ints(nums)

	// stop at nums[n-3], since the inner loop will use n-2 and n-1 when i == n-3
	for i := 0; i < n-2; i++ {

		// when nums[i] == nums[i-1], we encountered the same item we've just checked. skip this item.
		// the "i>0" protects us from null pointer since nums[0-1] is null pointer
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right, target := i+1, n-1, -nums[i]
		// we get an possible output when nums[left] + nums[right] + nums[i] == 0
		// -> nums[left] + nums[right] == -nums[i]
		// thus target is -nums[i]
		for left < right {
			if nums[left]+nums[right] == target {
				// found a solution, append it to the solutions
				output = append(output, []int{nums[i], nums[left], nums[right]})
				// if the next left is the same as current left, keep moving left
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// if the next right is the same as current right, keep moving right
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// now since we skipped all the duplicated left and right, we can move left and right to new values
				left++
				right--
			} else if nums[left]+nums[right] > target {
				// this case since sum > target, we want the sum to be smaller.
				// since nums is sorted, move right bounder to left would result in a smaller sum
				right--
			} else {
				// this case since sum < target, we want the sum to be bigger.
				// since nums is sorted, move left bounder to rig would result in a bigger sum
				left++
			}
		}
	}
	return output
}
