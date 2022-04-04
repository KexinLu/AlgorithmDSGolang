package sort

import (
	"math"
)

func radix(input []int) []int {
	maxD := findMaxDig(input)

	for i := 0; i <= maxD; i++ {
		input = cSort(input, i)
	}
	return input
}

func cSort(input []int, dig int) []int {
	divisor := int(math.Pow(10, float64(dig)))

	bucket := make([][]int, 10)
	for i := 0; i < len(input); i++ {
		cur := (input[i] / divisor) % 10
		bucket[cur] = append(bucket[cur], input[i])
	}
	result := []int{}
	for i := 0; i < 10; i++ {
		result = append(result, bucket[i]...)
	}

	return result
}

func findMaxDig(input []int) int {
	max := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}

	i := 1
	for {
		divisor := int(math.Pow(10, float64(i)))
		res := max / divisor
		if res == 0 {
			break
		}

		i++
	}

	return i
}
