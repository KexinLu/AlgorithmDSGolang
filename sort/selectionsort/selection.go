package sort

func selectionSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		minIndex := i
		for j := i; j < len(input); j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		input[minIndex], input[i] = input[i], input[minIndex]
	}
	return input
}
