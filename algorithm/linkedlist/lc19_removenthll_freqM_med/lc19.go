package lc19

type ListNode struct {
	Val  int
	Next *ListNode
}

// with array, cheating
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	leng := len(nodes)
	tgt := leng - n
	if tgt < 0 {
		return nodes[0]
	}

	if tgt == 0 {
		return nodes[1]
	}

	nodes[tgt-1].Next = nodes[tgt+1]
	return nodes[0]
}

func removeNthFromEndLL(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	ptr1 := head
	ptr2 := head
	for n > 0 {
		ptr2 = ptr2.Next
		n--
	}

	if ptr2 == nil {
		return ptr1.Next
	}
	for ptr2.Next != nil {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	ptr1.Next = ptr1.Next.Next
	return head
}
