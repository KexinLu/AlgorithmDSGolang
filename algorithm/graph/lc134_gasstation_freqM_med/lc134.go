package lc134

/**
 https://leetcode.com/problems/gas-station/
**/
// Brute Force
func canCompleteCircuitBF(gas []int, cost []int) int {
	gn := len(gas)
	for i := 0; i < gn; i++ {
		if gas[i] == 0 {
			continue
		}
		if gas[i] < cost[i] {
			continue
		}

		if check(i, gas, cost) {
			return i
		}
	}
	return -1
}

func check(start int, gas, cost []int) bool {
	gn := len(gas)
	gasTotal := 0
	i := start
	for {
		gasTotal += gas[i]
		if cost[i] <= gasTotal {
			gasTotal -= cost[i]
		} else {
			return false
		}

		if i < gn-1 {
			i++
		} else {
			i = 0
		}
		if i == start {
			break
		}
	}

	return true
}

// optimized O(N)
// if we start from i
// but at i+m we can't reach the next station
// everything between [i, i+m] doesn't work as a starting point
// reason: if we start at i+1, we either have leftover gas when travelling from i to i+1 or nothing left, we will never have a negative number (cuz we would just reset cur)
// thus having a previous station is always a better situation
// if the better situation doesn't work
// the worst situation won't work either
// thus we can just set the starting point at i+m+1
func canCompleteCircuit(gas []int, cost []int) int {
	gn := len(gas)
	totalGas := 0
	totalCost := 0
	cur := 0
	starting := 0
	for i := 0; i < gn; i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		cur += gas[i] - cost[i]
		if cur < 0 {
			cur = 0
			starting = i + 1
		}
	}

	if totalCost > totalGas {
		return -1
	}

	if starting > gn-1 {
		return -1
	}

	return starting
}
