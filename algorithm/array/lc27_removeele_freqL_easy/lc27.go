package lc27

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}
	p1 := 0
	p2 := 0
	count := 0
	for p2 < n {
		if nums[p2] == val {
			p2++
			continue
		}
		nums[p1] = nums[p2]
		count++
		p1++
		p2++
	}
	return count

}
