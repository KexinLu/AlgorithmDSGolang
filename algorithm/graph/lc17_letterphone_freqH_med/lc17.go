package lc17

// typical bfs traversal
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	return letterRecur(digits, 0, "")
}

func letterRecur(digits string, ptr int, carry string) []string {
	charMap := map[byte][]byte{
		'1': []byte{},
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
		'0': []byte{},
	}
	dn := len(digits)
	if ptr == dn {
		return []string{carry}
	}
	curChar := digits[ptr]

	if val, ok := charMap[curChar]; ok {
		// skip if 1 or 0
		if len(val) == 0 {
			return letterRecur(digits, ptr+1, carry)
		}

		rst := []string{}
		for _, v := range val {
			rst = append(rst, letterRecur(digits, ptr+1, carry+string(v))...)
		}

		return rst
	}

	return []string{}
}
