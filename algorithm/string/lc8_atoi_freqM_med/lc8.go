package lc8

import "math"

func myAtoi(s string) int {

	if len(s) == 0 {
		return 0
	}
	charMap := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'-': -1,
		'+': 1,
	}

	rst := 0
	sign := 1
	found := false
	for i := 0; i < len(s); i++ {
		cur := s[i]
		if found && (cur > '9' || cur < '0') {
			break
		}
		if cur == ' ' {
			continue
		}

		if cur == '+' || cur == '-' {
			found = true
			sign = charMap[cur]
			continue
		}

		if cur >= '0' && cur <= '9' {
			found = true
			rst = rst*10 + charMap[cur]
			if rst > math.MaxInt32 {
				break
			}
			continue
		}

		break

	}

	if sign > 0 {
		return min(rst, math.MaxInt32)
	}
	return max(sign*rst, math.MinInt32)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
