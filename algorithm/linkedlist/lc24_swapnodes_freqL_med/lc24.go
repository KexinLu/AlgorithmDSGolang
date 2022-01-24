package lc24

/**
* Definition for singly-linked list.
* type ListNode struct {
    *     Val int
    *     Next *ListNode
    * }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	root := &ListNode{
		Next: head,
	}
	ans := root
	node := head
	for node != nil && node.Next != nil {
		ans.Next = node.Next
		next := node.Next
		tmp := next.Next
		next.Next = node
		node.Next = tmp
		ans = node
		node = node.Next
	}
	return root.Next
}
