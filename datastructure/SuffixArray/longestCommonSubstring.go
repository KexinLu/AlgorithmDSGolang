package suffixArray

import (
	"bytes"
	"sort"
)

func CreateSuffixArray(s string) []string {
	suffixArray := make([]string, 0)

	for i, _ := range s {
		suffixArray = append(suffixArray, s[i:])
	}

	sort.Strings(suffixArray)

	return suffixArray
}

// longest common prefix array
func CreateLCPArray(arr []string) (rst []int) {
	rst = make([]int, 0)

	if len(arr) < 1 {
		return rst
	}

	rst = append(rst, 0)
	if len(arr) < 2 {
		return rst
	}

	for i, j := 0, 1; j <= len(arr)-1; i, j = i+1, j+1 {
		common := 0
		for idx, a := range arr[i] {
			if a == rune(arr[j][idx]) {
				common++
				continue
			} else {
				break
			}
		}
		rst = append(rst, common)
	}

	return rst
}

func getCommon(s1 string, s2 string) string {
	var buffer bytes.Buffer
	for idx, a := range s1 {
		char := s2[idx]
		if a == rune(char) {
			buffer.WriteString(string(char))
			continue
		} else {
			break
		}
	}

	return buffer.String()
}

func FindLongestRepeatedSubString(s string) []string {
	suffArr := CreateSuffixArray(s)
	lcp := CreateLCPArray(suffArr)

	maxLen := 0
	strs := make([]string, 0)
	for idx, v := range lcp {
		if v == maxLen && v != 0 {
			strs = append(strs, getCommon(suffArr[idx-1], suffArr[idx]))
		} else if v > maxLen {
			maxLen = v
			strs = []string{
				getCommon(suffArr[idx-1], suffArr[idx]),
			}
		}
		continue
	}

	return strs
}
