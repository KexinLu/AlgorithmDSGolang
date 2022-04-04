package sort

func insert(input []int) []int {
	for i := 1; i < len(input); i++ {
		for j := i; j > 0; j-- {
			if input[j] < input[j-1] {
				input[j], input[j-1] = input[j-1], input[j]
			}
		}
	}

	return input
}
