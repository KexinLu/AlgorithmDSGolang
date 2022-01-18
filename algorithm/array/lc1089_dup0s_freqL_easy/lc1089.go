package lc1089

/**
https://leetcode.com/problems/duplicate-zeros/

Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.



Example 1:

Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
Example 2:

Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]


Constraints:

1 <= arr.length <= 104
0 <= arr[i] <= 9


**/

// naive solution
func duplicateZeros(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			for j := n - 1; j > i; j-- {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			arr[i] = 0
			i++
		}
	}
}

// concat solution
func dupZeros(arr []int) {
	n := len(arr)
	skip := false

	for i := 0; i < n; i++ {

		tmp := []int{}

		if skip {
			skip = false
			continue
		}
		if arr[i] == 0 {
			tmp = append(arr[:i], 0)
			arr = append(tmp, arr[i:n-1]...)
			skip = true
		}
	}
}

func dupZeroAlgorithm(arr []int) {
	// counting how many zeros we have
	count, n := 0, len(arr)

	// first loop thorugh arr once, O(n)
	for _, v := range arr {
		if v == 0 {
			count++
		}
	}

	// loop through arr one more time, O(n)
	for i := n - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if count+i < n {
				arr[count+i] = 0
			}

			if count-1+i < n {
				arr[count-1+i] = 0
			}

			count--
		} else if i+count < n {
			arr[count+i] = arr[i]
		}
	}
}
