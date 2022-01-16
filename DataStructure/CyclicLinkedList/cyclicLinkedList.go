package cyclicLinkedList

import (
	"bytes"
	"fmt"
)

type LinkedList struct {
	length int
	head   *Node
}

type Node struct {
	next  *Node
	prev  *Node
	value int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) AddToHead(val int) {
	if ll.isEmpty() {
		ll.addFirst(val)
		return
	}

	curHead := ll.head
	prev := curHead.prev
	newNode := &Node{
		prev:  prev,
		next:  curHead,
		value: val,
	}
	ll.head = newNode
	ll.length++
	return
}

func (ll *LinkedList) AddToTail(val int) {
	if ll.isEmpty() {
		ll.addFirst(val)
		return
	}

	curHead := ll.head
	curTail := curHead.prev
	n := &Node{
		prev:  curTail,
		next:  curHead,
		value: val,
	}
	curHead.prev = n
	curTail.next = n
	ll.length++
}

func (ll *LinkedList) addFirst(val int) {
	n := &Node{
		value: val,
	}
	n.prev = n
	n.next = n
	ll.head = n
	ll.length++
	return
}

func (ll *LinkedList) Reverse() {
	var nxt *Node
	tail := ll.head.prev
	for cur, i := ll.head, 0; i < ll.length; i += 1 {
		nxt = cur.next
		cur.next = cur.prev
		cur.prev = nxt
		cur = nxt
	}
	ll.head = tail
}

func (ll *LinkedList) ToString() string {
	var buffer bytes.Buffer
	if ll.isEmpty() {
		return ""
	}

	for cur, i := ll.head, 0; i < ll.length; i++ {
		buffer.WriteString(fmt.Sprint(cur.value))
		if i < ll.length-1 {
			buffer.WriteString(",")
		}
		cur = cur.next
	}

	return buffer.String()
}

func (ll *LinkedList) ToStringDetail() string {
	var buffer bytes.Buffer
	if ll.isEmpty() {
		return ""
	}

	for cur, i := ll.head, 0; i < ll.length; i++ {
		buffer.WriteString(fmt.Sprintf("{%d, %d, %d}", cur.prev.value, cur.value, cur.next.value))
		if i < ll.length-1 {
			buffer.WriteString(",")
		}
		cur = cur.next
	}

	return buffer.String()
}

func (ll *LinkedList) isEmpty() bool {
	return ll.head == nil
}
