package lc28

// https://leetcode.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {

	hn, nn := len(haystack), len(needle)
	for i := 0; i < hn-nn+1; i++ {
		j := 0
		m := i
		for m < hn && j < nn && haystack[m] == needle[j] {
			m++
			j++
		}
		if j >= nn {
			return i
		}
	}
	return -1
}
