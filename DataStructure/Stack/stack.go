package stack

import "fmt"

// FILO
type Stack struct {
	top   *Node
	count int
}

// [] [] [] []

type Node struct {
	value int
	next  *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	n := &Node{
		value: val,
		next:  s.top,
	}

	s.top = n
	s.count++
}

func (s *Stack) Pop() (int, bool) {
	rst := s.top
	if rst == nil {
		return 0, false
	}
	nxt := rst.next
	s.top = nxt
	rst.next = nil
	s.count--

	return rst.value, true
}

func (s *Stack) Print() {
	tailWalk(s.top, func(n *Node) {
		fmt.Println(n.value)
	})
}

func tailWalk(n *Node, fn func(n *Node)) {
	if n == nil {
		return
	}

	if n.next == nil {
		fn(n)
		return
	}

	tailWalk(n.next, fn)
	fn(n)
}

func (s *Stack) Peak() int {
	return s.top.value
}

func (s *Stack) isEmpty() bool {
	return s.count == 0
}
