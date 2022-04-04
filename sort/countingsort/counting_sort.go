package sort

func countingSort(input []int) []int {
	max := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}

	bucket := make([][]int, max+1)

	for i := 0; i < len(input); i++ {
		bucket[input[i]] = append(bucket[input[i]], input[i])
	}

	result := []int{}
	for i := 0; i < len(bucket); i++ {
		result = append(result, bucket[i]...)
	}

	return result
}
