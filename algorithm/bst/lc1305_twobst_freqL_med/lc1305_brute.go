package lc1305

/**
* Definition for a binary tree node.
**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// simple walk both then merge
func getAllElementsBrute(root1 *TreeNode, root2 *TreeNode) []int {
	a1 := []int{}
	a2 := []int{}
	walk(root1, &a1)
	walk(root2, &a2)
	i := 0
	j := 0
	out := make([]int, len(a1)+len(a2))
	count := 0
	for count < len(a1)+len(a2) {
		if i >= len(a1) {
			for j < len(a2) && count < len(a1)+len(a2) {
				out[count] = a2[j]
				j++
				count++
			}
			break
		}
		if j >= len(a2) {
			for i < len(a1) && count < len(a1)+len(a2) {
				out[count] = a1[i]
				i++
				count++
			}
			break
		}
		if a1[i] <= a2[j] {
			out[count] = a1[i]
			i++
		} else {
			out[count] = a2[j]
			j++
		}
		count++
	}
	return out
}

func walk(t *TreeNode, result *[]int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		walk(t.Left, result)
	}
	*result = append(*result, t.Val)
	walk(t.Right, result)
}
