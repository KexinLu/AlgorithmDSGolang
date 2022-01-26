package lc1305

type Stack []*TreeNode

func (stk *Stack) Pop() *TreeNode {
	out := (*stk)[len(*stk)-1]
	*stk = (*stk)[:len(*stk)-1]
	return out
}

func (stk *Stack) Push(i *TreeNode) {
	*stk = append(*stk, i)
}

func (stk Stack) Peek() int {
	return stk[len(stk)-1].Val
}
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	out := []int{}
	s1 := make(Stack, 0)
	s2 := make(Stack, 0)
	for {
		if len(s1) <= 0 && len(s2) <= 0 && root1 == nil && root2 == nil {
			break
		}
		for root1 != nil {
			s1.Push(root1)
			root1 = root1.Left
		}

		for root2 != nil {
			s2.Push(root2)
			root2 = root2.Left
		}

		if len(s2) == 0 {
			root1 = s1.Pop()
			out = append(out, root1.Val)
			root1 = root1.Right
		} else if len(s1) == 0 {
			root2 = s2.Pop()
			out = append(out, root2.Val)
			root2 = root2.Right
		} else if s1.Peek() <= s2.Peek() {
			root1 = s1.Pop()
			out = append(out, root1.Val)
			root1 = root1.Right
		} else {
			root2 = s2.Pop()
			out = append(out, root2.Val)
			root2 = root2.Right
		}
	}
	return out
}
