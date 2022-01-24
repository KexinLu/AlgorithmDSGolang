package lc23heap

/**
* Definition for singly-linked list.
* type ListNode struct {
    *     Val int
    *     Next *ListNode
    * }
*/

/**
 * Definition for singly-linked list.
 */

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type PQ []*ListNode

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//!important this is avoiding memory leak
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	if pq[i].Val < pq[j].Val {
		return true
	}
	return false
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PQ, 0)
	for _, l := range lists {
		for l != nil {
			next := l.Next
			l.Next = nil
			heap.Push(&pq, l)
			l = next
		}
	}
	root := &ListNode{}
	node := root

	for len(pq) != 0 {
		cur := heap.Pop(&pq)
		node.Next = cur.(*ListNode)
		node = node.Next

	}
	return root.Next
}
