package SingleLinkedList

import (
	"bytes"
	"fmt"
)

type SingleLinkedList struct {
	length int
	head   *Node
}

type Node struct {
	val  int
	next *Node
}

type ILinkedList interface {
	AddNodeToHead(val int)
	AddNodeToTail(val int)
	GetHead() *Node
	RemoveHead()
	RemoveTail()
	Reverse()
	Count() int
	ReversePartition(start int, end int) error
	Display()
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (l *SingleLinkedList) GetHead() *Node {
	return l.head
}

func (l *SingleLinkedList) RemoveHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

func (l *SingleLinkedList) RemoveTail() {
	if l.head == nil {
		return
	} else if l.head.next == nil {
		l.head = nil
		return
	}

	prev := l.head
	cur := prev.next
	for cur.next != nil {
		prev = cur
		cur = cur.next
	}

	prev.next = nil
}

func (l *SingleLinkedList) AddToHead(val int) {
	cur := l.head
	if cur == nil {
		l.head = &Node{
			val: val,
		}
	}

	next := cur
	l.head = &Node{
		val:  val,
		next: next,
	}
	l.length += 1
}

func (l *SingleLinkedList) AddToTail(val int) {
	cur := l.head
	if cur == nil {
		l.head = &Node{
			val: val,
		}

		l.length += 1
		return
	} else if cur.next == nil {
		cur.next = &Node{
			val: val,
		}

		l.length += 1
		return
	}

	for cur.next != nil {
		cur = cur.next
	}

	cur.next = &Node{
		val: val,
	}
	l.length += 1

}

func (l *SingleLinkedList) Display() {
	for cur := l.head; cur != nil; cur = cur.next {
		fmt.Printf("%v,", cur.val)
	}

	fmt.Print("\n")
}

func (l *SingleLinkedList) ToString() string {
	var buffer bytes.Buffer
	for cur := l.head; cur != nil; cur = cur.next {
		buffer.WriteString(fmt.Sprint(cur.val))
		if cur.next != nil {
			buffer.WriteString(",")
		}
	}

	return buffer.String()
}

func (l *SingleLinkedList) Reverse() {
	var prev, next *Node

	cur := l.head
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	l.head = prev
}
