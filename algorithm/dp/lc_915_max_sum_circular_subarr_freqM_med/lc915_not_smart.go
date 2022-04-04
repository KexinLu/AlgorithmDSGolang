package lc915

import "math"

type Q []int

func (q Q) Top() int {
	if len(q) == 0 {
		return -1
	}
	return q[0]
}

func (q *Q) Pop() int {
	if len(*q) == 0 {
		return -1
	}

	first := (*q)[0]
	(*q) = (*q)[1:]

	return first
}

func (q *Q) Append(in int, p []int) {
	cur := p[in]
	upperBound := len(*q)
	for i := len(*q) - 1; i >= 0; i-- {
		item := p[(*q)[i]]
		if item >= cur {
			upperBound = i
		} else {
			break
		}
	}

	(*q) = append((*q)[:upperBound], in)
}

func maxSubarraySumCircular(nums []int) int {
	limit := len(nums)
	p := make([]int, 2*limit)
	p[0] = 0
	for i := 1; i < 2*limit; i++ {
		p[i] = p[i-1] + nums[(i-1)%limit]
	}
	monoQueue := Q{}
	res := math.MinInt32
	top := monoQueue.Top()
	for i := 0; i < 2*limit; i++ {
		for top != -1 && i-top >= limit {
			monoQueue.Pop()
			top = monoQueue.Top()
		}

		if top != -1 {
			best := nums[i%limit]
			if p[i]-p[top] > 0 {
				best += p[i] - p[top]
			}
			res = maxOf(best, res)
		} else {
			best := nums[i%limit]
			if p[i] > 0 {
				best += p[i]
			}
			res = maxOf(res, best)
		}
		monoQueue.Append(i, p)
		top = monoQueue.Top()
	}

	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}

	return b
}
