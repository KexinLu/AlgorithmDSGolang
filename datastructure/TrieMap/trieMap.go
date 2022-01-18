package trie

import (
	"bytes"
	"fmt"
)

type Node struct {
	children map[rune]*Node
	isLeaf   bool
}

func NewTrie() *Node {
	return &Node{
		children: make(map[rune]*Node, 0),
		isLeaf:   false,
	}
}

func (n *Node) IsLeaf() bool {
	return n.isLeaf
}

func (n *Node) Insert(s string) {
	// a b c d e
	var cur *Node
	cur = n
	for _, c := range s {
		next, ok := n.children[c]
		if !ok {
			next = NewTrie()
			cur.children[c] = next
		}

		cur = next
	}

	cur.isLeaf = true
}

func (n *Node) Find(s string) bool {
	next, ok := n, false
	for _, c := range s {
		next, ok = next.children[c]
		if !ok {
			return false
		}
	}
	return next.IsLeaf()
}

func (n *Node) Size() int {
	r := 0
	for _, c := range n.children {
		r += c.Size()
	}

	if n.isLeaf() {
		r++
	}

	return r
}

func (n *Node) PrintByLayer() {
	toPrint := make([]*Node, 0)
	var buffer bytes.Buffer
	for k, v := range n.children {
		buffer.WriteString(fmt.Sprintf(" %v ", string(k)))
		toPrint = append(toPrint, v)
	}
	fmt.Println(buffer.String())

	for _, toP := range toPrint {
		toP.PrintByLayer()
	}
}
