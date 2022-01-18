package binarySearchTree

import (
	"bytes"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BSTree struct {
	root  *Node
	count int
}

func NewBST() *BSTree {
	return &BSTree{}
}

func (bst *BSTree) Add(val int) {
	if bst.root == nil {
		bst.root = &Node{
			value: val,
		}
		bst.count++

		return
	}

	insert(bst.root, val)
	bst.count++
}

func insert(n *Node, val int) *Node {
	if n == nil {
		return &Node{
			value: val,
		}
	}

	if val < n.value {
		n.left = insert(n.left, val)
	} else {
		n.right = insert(n.right, val)
	}

	return n
}

func walkInOrderRecursive(n *Node, fn func(*Node)) {
	if n == nil {
		return
	}

	walkInOrderRecursive(n.left, fn)
	fn(n)
	walkInOrderRecursive(n.right, fn)
}

func walkInOrder(n *Node, fn func(*Node)) {
	walkInOrderRecursive(n, fn)
}

func walkPreOrder(n *Node, fn func(*Node)) {
	walkPreOrderRecursive(n, fn)
}

// root left right
func walkPreOrderRecursive(n *Node, fn func(*Node)) {
	if n == nil {
		return
	}

	fn(n)
	walkPreOrderRecursive(n.left, fn)
	walkPreOrderRecursive(n.right, fn)
}

func walkPostOrder(n *Node, fn func(*Node)) {
	walkPostOrderRecursive(n, fn)
}

// root left right
func walkPostOrderRecursive(n *Node, fn func(*Node)) {
	if n == nil {
		return
	}

	walkPostOrderRecursive(n.left, fn)
	walkPostOrderRecursive(n.right, fn)
	fn(n)
}

func (bst *BSTree) Print() {
	walkInOrder(bst.root, func(n *Node) {
		fmt.Println(n.value)
	})
}

func (bst *BSTree) ToStringInOrder() string {
	var buffer bytes.Buffer
	walkInOrder(bst.root, func(n *Node) {
		buffer.WriteString(fmt.Sprintf(" %v ", n.value))
	})
	return buffer.String()
}

func (bst *BSTree) ToStringPreOrder() string {
	var buffer bytes.Buffer
	walkPreOrder(bst.root, func(n *Node) {
		buffer.WriteString(fmt.Sprintf(" %v ", n.value))
	})

	return buffer.String()
}

func (bst *BSTree) ToStringPostOrder() string {
	var buffer bytes.Buffer
	walkPostOrder(bst.root, func(n *Node) {
		buffer.WriteString(fmt.Sprintf(" %v ", n.value))
	})
	return buffer.String()
}

func (bst *BSTree) walkByLevel(fn func(*Node)) {
	if bst.root == nil {
		return
	}

	var q []*Node
	var n *Node

	q = append(q, bst.root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		fn(n)
		if n.left != nil {
			q = append(q, n.left)
		}

		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

func (bst *BSTree) ToStringByLevel() {
	bst.walkByLevel(func(node *Node) {
		fmt.Println(node.value)
	})
}
