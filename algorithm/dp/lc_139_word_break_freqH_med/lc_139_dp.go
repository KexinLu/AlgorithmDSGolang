package lc_139

func wordBreakDP(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	wordMap := make(map[string]int)
	for _, val := range wordDict {
		wordMap[val] = 0
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			cur := dp[j]
			if cur {

				if _, ok := wordMap[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(s)]

}
