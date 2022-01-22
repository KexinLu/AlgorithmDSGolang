package lc20

/**
https://leetcode.com/problems/valid-parentheses/
**/
type Stack []byte

func (s *Stack) Push(b byte) {
	*s = append(*s, b)
}

func (s *Stack) Pop() byte {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func (s Stack) Empty() bool {
	if len(s) <= 0 {
		return true
	}
	return false
}

func isValid(s string) bool {
	if len(s)&1 == 1 {
		return false
	}
	charMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	revCharMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stk := make(Stack, 0)

	for i := 0; i < len(s); i++ {
		if _, ok := charMap[s[i]]; ok {
			stk.Push(s[i])
		} else {
			if stk.Empty() {
				return false
			}
			top := stk.Pop()
			if revCharMap[s[i]] != top {
				return false
			}
		}
	}

	if !stk.Empty() {
		return false
	}
	return true

}

func isValidSimple(s string) bool {
	charMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stk := []byte{}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			stk = append(stk, s[i])
		case ')', '}', ']':
			if len(stk) == 0 {
				return false
			}
			top := stk[len(stk)-1]
			if s[i] != charMap[top] {
				return false
			}
			stk = stk[:len(stk)-1]
		}
	}

	return len(stk) == 0
}
