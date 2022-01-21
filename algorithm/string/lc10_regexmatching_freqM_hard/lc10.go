package lc10

// simple recur
func isMatch(s string, p string) bool {
	pn, sn := len(p), len(s)
	if pn == 0 {
		// no more pattern to check

		// if there's still string to check it's false
		return sn == 0
	}

	// pn > 0
	// if sn == 0 then no match
	firstMatch := sn != 0 && (p[0] == s[0] || p[0] == '.')

	// forward lookup
	if pn >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || firstMatch && isMatch(s[1:], p)
	}

	return firstMatch && isMatch(s[1:], p[1:])

}

// recur with memo

func isMatchMemo(s string, p string) bool {
	pn, sn := len(p), len(s)

	// initialize the memo
	memo := make([][]int, sn+1, sn+1)
	for idx, _ := range memo {
		memo[idx] = make([]int, pn+1, pn+1)
	}

	return isMatchHelper(0, 0, s, p, memo)
}

func isMatchHelper(i, j int, s, p string, memo [][]int) (rst bool) {
	pn, sn := len(p), len(s)
	if memo[i][j] != 0 {
		return memo[i][j] > 0
	}

	if j == pn {
		rst = i == sn
	} else {
		// not done
		aMatch := i < sn && (p[j] == s[i] || p[j] == '.')

		if j+1 < pn && p[j+1] == '*' {
			return isMatchHelper(i, j+2, s, p, memo) || aMatch && isMatchHelper(i+1, j, s, p, memo)
		}

		return aMatch && isMatchHelper(i+1, j+1, s, p, memo)

	}

	if rst {
		memo[i][j] = 1
	} else {
		memo[i][j] = -1
	}

	return
}
