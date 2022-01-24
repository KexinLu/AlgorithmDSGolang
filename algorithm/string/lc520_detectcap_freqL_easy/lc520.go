package lc520

func detectCapitalUse(word string) bool {
	n := len(word)
	isCap := false
	if word[0] >= 'A' && word[0] <= 'Z' {
		isCap = true
	}

	capCount := 0
	for i := 1; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			if !isCap {
				return false
			}
			capCount++
		}
	}

	if isCap && (capCount == n-1 || capCount == 0) {
		return true
	}

	if !isCap {
		return capCount == 0
	}
	return false
}
