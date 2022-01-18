package BTree

type Node struct {
	value  int
	parent *Node
	left   *Node
	right  *Node
}

type BTree struct {
	root *Node
	size int
}

func CreateBTree() *BTree {
	return &BTree{
		size: 0,
	}
}

func Insert(n *Node, parent *Node, val int) *Node {
	if n == nil {
		return &Node{
			value:  val,
			parent: parent,
		}
	}

	if n <= n.value {
		n.left = Insert(n.left, n, val)
	} else {
		n.right = Insert(n.right, n, val)
	}
	return n
}

func Delete(n *Node, val int) *Node {
	if n == nil {
		return nil
	}

	if val < n.value {
		n.left = Delete(n.left, val)
	} else if val > n.value {
		n.right = Delete(n.right, val)
	} else {
		//deleting this node
		if n.right == nil {
			return n.left
		}
		if n.left == nil {
			return n.right
		}
		// both side has child
		s := findSuccessor(n.right)
		s.left = n.left

		return n.right
	}
	return n
}
