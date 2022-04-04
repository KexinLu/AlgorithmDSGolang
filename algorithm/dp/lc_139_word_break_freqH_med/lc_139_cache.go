package lc139

//https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	wordMap := make(map[string]int)
	for _, val := range wordDict {
		wordMap[val] = 0
	}

	cache := make(map[int]bool)
	return wordBreakRecur(s, 0, wordMap, &cache)

}

func wordBreakRecur(s string, start int, wordMap map[string]int, cache *map[int]bool) bool {
	if len(s[start:]) == 0 {
		return true
	}

	if _, ok := (*cache)[start]; ok {
		return (*cache)[start]
	}

	for e := start + 1; e <= len(s); e++ {
		if _, ok := wordMap[s[start:e]]; ok {
			if wordBreakRecur(s, e, wordMap, cache) {
				(*cache)[start] = true
				return true
			}
		}
	}

	(*cache)[start] = false
	return false
}
