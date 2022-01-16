package DoubleLinkedList

import (
	"bytes"
	"fmt"
)

type DoubleLinkedList struct {
	length int
	head   *Node
}

type Node struct {
	next *Node
	prev *Node
	val  int
}

func NewDouble() (dbl *DoubleLinkedList) {
	return &DoubleLinkedList{}
}

func (dbl *DoubleLinkedList) AddToHead(val int) {
	cur := dbl.head
	if cur == nil {
		dbl.head = &Node{
			val: val,
		}
		dbl.length += 1
		return
	}

	nd := &Node{
		next: cur,
		val:  val,
	}

	cur.prev = nd
	dbl.head = nd
	dbl.length += 1
}

func (dbl *DoubleLinkedList) AddToTail(val int) {
	cur := dbl.head
	if cur == nil {
		dbl.length += 1
		dbl.head = &Node{
			val: val,
		}
		return
	}

	var prev *Node
	for cur != nil {
		prev = cur
		cur = cur.next
	}

	prev.next = &Node{
		val:  val,
		prev: prev,
	}
	dbl.length += 1
}

func (dbl *DoubleLinkedList) ToString() string {
	var buffer bytes.Buffer

	nd := dbl.head
	for nd != nil {
		buffer.WriteString(fmt.Sprint(nd.val))
		if nd.next != nil {
			buffer.WriteString(",")
		}
		nd = nd.next
	}

	return buffer.String()
}

func (dbl *DoubleLinkedList) DeleteHead() {
	head := dbl.head
	if head == nil {
		return
	} else if head.next == nil {
		dbl.head = nil
		dbl.length = 0
	}
	head.next.prev = nil
	dbl.head = head.next
	head.next = nil
}

func (dbl *DoubleLinkedList) DeleteTail() {
	head := dbl.head
	if head == nil {
		return
	} else if head.next == nil {
		dbl.head = nil
		dbl.length = 0
	}

	var prev *Node
	cur := dbl.head
	for cur.next != nil {
		prev = cur
		cur = cur.next
	}

	prev.next = nil
	cur.prev = nil
	dbl.length -= 1
}
