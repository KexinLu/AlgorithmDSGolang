package lc21

/**
 * Definition for singly-linked list.
  * type ListNode struct {
       *     Val int
        *     Next *ListNode
	 * }
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list2
	rst := &ListNode{
		Val:  -1,
		Next: nil,
	}
	head := rst

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			head.Next = p1
			p1 = p1.Next
		} else {
			head.Next = p2
			p2 = p2.Next
		}
		head = head.Next
	}

	if p1 != nil {
		head.Next = p1
	}
	if p2 != nil {
		head.Next = p2
	}
	return rst.Next
}
