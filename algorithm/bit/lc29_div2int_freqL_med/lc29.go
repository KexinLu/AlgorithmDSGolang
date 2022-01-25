package lc29

import (
	"math"
)

// https://leetcode.com/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	//
	if dividend == 0 {
		return 0
	}

	ds := dividend >> 63
	dvs := divisor >> 63
	sign := ds ^ dvs

	if ds == -1 {
		dividend = (dividend ^ -1) + 1
	}
	if dvs == -1 {
		divisor = (divisor ^ -1) + 1
	}
	org := divisor
	carry := 0
	count := 0
	flag := 1
	for carry < dividend {
		carry += divisor
		count += flag
		divisor <<= 1
		flag <<= 1
	}

	for int(math.Abs(float64(carry-dividend))) >= org {
		divisor >>= 1
		flag >>= 1
		if carry-dividend > 0 {
			carry -= divisor
			count -= flag
		} else if carry-dividend < 0 {
			carry += divisor
			count += flag
		} else {
			break
		}

	}
	if carry-dividend > 0 {
		count--
	}
	if sign != 0 {
		return max((count^-1)+1, math.MinInt32)
	}
	return min(count, math.MaxInt32)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
