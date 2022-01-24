package lc23

import "math"

/**
 Definition for singly-linked list.
 type ListNode struct {
        Val int
        Next *ListNode
 }
	 https://leetcode.com/problems/merge-k-sorted-lists/
*/
type PQ []*ListNode

func (pq PQ) Push(x interface{}) {
	n := len(pq)
	pq = append(pq, x.(*ListNode))
}

func (pq PQ) Pop() interface{} {
	first := pq[0]
	pq = pq[1:]
	return first
}

func mergeKLists(lists []*ListNode) *ListNode {
	root := &ListNode{
		Val: math.MaxInt64,
	}
	cur := root
	for len(lists) > 0 {
		min := root
		minIdx := -1
		for idx, p := range lists {
			if p == nil {
				minIdx = idx
				min = p
				break
			}
			if p.Val < min.Val {
				min = p
				minIdx = idx
			}
		}
		if minIdx == -1 {
			break
		}
		if lists[minIdx] == nil {
			lists = append(lists[:minIdx], lists[minIdx+1:]...)
		}
		if min != nil {
			lists[minIdx] = min.Next
			cur.Next = min
			cur = cur.Next
		}
	}
	return root.Next
}
