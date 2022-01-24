package lc23heap

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {

	a3 := &ListNode{
		Val:  5,
		Next: nil,
	}
	a2 := &ListNode{
		Val:  4,
		Next: a3,
	}
	a1 := &ListNode{
		Val:  1,
		Next: a2,
	}

	b3 := &ListNode{
		Val:  4,
		Next: nil,
	}

	b2 := &ListNode{
		Val:  3,
		Next: b3,
	}

	b1 := &ListNode{
		Val:  1,
		Next: b2,
	}

	c2 := &ListNode{
		Val:  6,
		Next: nil,
	}
	c1 := &ListNode{
		Val:  2,
		Next: c2,
	}

	input := []*ListNode{a1, b1, c1}

	root := mergeKLists(input)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
}
