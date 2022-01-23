package lc1510

import "math"

// https://leetcode.com/problems/stone-game-iv/
func winnerSquareGameDFS(n int) bool {
	if n == 1 {
		return true
	}
	marker := make(map[int]bool)
	marker[0] = false
	return check(n, marker)
}

func check(n int, marker map[int]bool) bool {
	if prev, ok := marker[n]; ok {
		return prev
	}
	sqrt := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrt; i++ {
		if !check(n-i*i, marker) {
			marker[n-i*i] = false
			return true
		} else {
			marker[n-i*i] = true
		}
	}

	return false
}

func winnerSquareGameDP(n int) bool {
	marker := make([]bool, n+1)
	marker[0] = false
	for i := 1; i <= n; i++ {
		sq := int(math.Sqrt(float64(i)))
		if sq == i {
			marker[i] = true
		}
		for k := 1; k <= sq; k++ {
			if marker[i-k*k] == false {
				marker[i] = true
			}
		}
	}

	return marker[n]
}
