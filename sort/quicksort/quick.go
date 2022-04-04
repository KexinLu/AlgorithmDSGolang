package sort

func quickSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	///quickSortRecur(&input, 0, len(input)-1)
	quickSort2(&input, 0, len(input)-1)
	return input
}

func quickSortRecur(input *[]int, start, end int) {
	arr := *input
	if start >= end {
		return
	}
	pivot := end
	i := start
	j := start

	for i <= j && j < end {
		if arr[j] <= arr[pivot] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		j++
	}

	arr[i], arr[pivot] = arr[pivot], arr[i]
	pivot = i

	quickSortRecur(input, start, pivot-1)
	quickSortRecur(input, pivot+1, end)
}

func quickSort2(input *[]int, start, end int) {
	if start >= end {
		return
	}
	arr := *input

	i := start
	j := end - 1

	for i <= j && i < end && j >= start {
		if arr[i] > arr[end] && arr[j] <= arr[end] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
		for i <= j && i < end && arr[i] <= arr[end] {
			i++
		}
		for i <= j && j >= start && arr[j] >= arr[end] {
			j--
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	quickSort2(input, start, i-1)
	quickSort2(input, i+1, end)
}
