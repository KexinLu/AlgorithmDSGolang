package lc29

import (
	"math"
)

// https://leetcode.com/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		// MaxInt32 is to limit the result in -2^32 < result < 2^32-1
		return min(dividend, math.MaxInt32)
	}

	// since we have int64
	// bit move to the right 63 times
	// if you have a positive(include 0) >> 63 yields 0
	// if you have a negative number >> 63 yields -1

	// getting dividend's sign
	ds := dividend >> 63
	// getting divisor's sign
	dvs := divisor >> 63

	// using XOr to decide if we should flip the final result
	sign := ds ^ dvs

	// turn both dividend and divisor into positive number
	if ds == -1 {
		// since we can't use multiplication, technically -dividend can't be used (not sure)
		// thus use XOr to negate the number to negative - 1, then +1
		// this is because eg. -15 ^ -1 = 14, thus making it 15 we need to plus 1
		dividend = (dividend ^ -1) + 1
	}
	if dvs == -1 {
		divisor = (divisor ^ -1) + 1
	}

	// keep the original divisor
	org := divisor

	carry := 0
	count := 0

	// incr is how much we increment each time
	incr := 1

	// Algorithm :
	// if we have a huge number, increment by divisor will consume too much time, eg when divisor = 1 and dividend is maxInt64
	// thus we want to first exceed the dividend by doing
	// 1 + 2 + 4 + 8 + 16 + 32 + 64 + 128 + 256 +... until we exceed the dividend

	// then we shrink the increment by /2 each time
	// eg. -256, + 128, - 64, + 32, - 16, + 8 until we have Abs(carry - divident) < original divisor

	for carry < dividend {
		carry += divisor
		count += incr
		divisor <<= 1
		incr <<= 1
	}

	for int(math.Abs(float64(carry-dividend))) >= org {
		divisor >>= 1
		incr >>= 1
		if carry-dividend > 0 {
			carry -= divisor
			count -= incr
		} else if carry-dividend < 0 {
			carry += divisor
			count += incr
		} else {
			break
		}

	}
	// we want floor, thus this
	if carry-dividend > 0 {
		count--
	}

	// if one negative on positive, flip the result
	// limit the result within int32 boundary
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
