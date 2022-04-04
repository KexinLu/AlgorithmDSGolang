package sort

func bubble(input []int) []int {
	for i := len(input) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}
