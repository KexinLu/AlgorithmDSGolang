package lc19

import (
	"fmt"
	"testing"
	//. "github.com/smartystreets/goconvey/convey"
)

func TestRemove(t *testing.T) {
	var prev *ListNode
	var head *ListNode
	for _, v := range []int{1, 2, 3, 4, 5} {
		if prev == nil {
			prev = &ListNode{
				Val:  v,
				Next: nil,
			}
			head = prev
			continue
		}

		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		prev.Next = node
		prev = node
	}

	n := head
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
	removeNthFromEndLL(head, 3)
	n = head
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}

}
