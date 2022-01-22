package lc12

// optimized
// Greedy algorithm
func intToRoman(num int) string {
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	index, res := 0, ""
	for num > 0 {
		if num-value[index] >= 0 {
			res += symbol[index]
			num -= value[index]
		} else {
			index++
		}
	}
	return res
}
