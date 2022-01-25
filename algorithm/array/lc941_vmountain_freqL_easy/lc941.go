package lc941

func validMountainArray(arr []int) bool {
	n := len(arr)
	if len(arr) < 3 {
		return false
	}

	i := 0
	for i+1 < n && arr[i] < arr[i+1] {
		if arr[i] >= arr[i+1] {
			return false
		}
		i++
	}
	if i == 0 || i == n-1 {
		return false
	}

	for i+1 < n {
		if arr[i+1] >= arr[i] {
			return false
		}
		i++
	}
	if i+1 != n {
		return false
	}
	return true
}
