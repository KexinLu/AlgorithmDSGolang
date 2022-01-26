package lc31

/**
https://leetcode.com/problems/next-permutation/
**/

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	if n == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}
	first := nums[0]
	next := 0
	for i := n - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			cur := nums[i-1]
			j := i
			for j < len(nums) && nums[j] > cur {
				j++
			}
			nums[j-1], nums[i-1] = nums[i-1], nums[j-1]
			reverse(&nums, i, len(nums)-1)
			return
		}
		if next == 0 && nums[i] > first {
			next = i
		}
	}
	if next == 0 {
		reverse(&nums, 0, len(nums)-1)
		return
	}
	nums[next], nums[0] = nums[0], nums[next]
	reverse(&nums, 1, len(nums)-1)
}

func reverse(nums *[]int, i, j int) {
	arr := *nums
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
