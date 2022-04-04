package sort

func mergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	//mid := len(input) >> 1
	mid := len(input) / 2

	r1 := mergeSort(input[:mid])
	r2 := mergeSort(input[mid:])

	return merge(r1, r2)
}

func merge(in1, in2 []int) []int {
	i, j := 0, 0
	result := []int{}
	for i < len(in1) && j < len(in2) {
		if in1[i] <= in2[j] {
			result = append(result, in1[i])
			i++
		} else {
			result = append(result, in2[j])
			j++
		}
	}

	if i >= len(in1) {
		result = append(result, in2[j:]...)
	} else {
		result = append(result, in1[i:]...)
	}

	return result
}
