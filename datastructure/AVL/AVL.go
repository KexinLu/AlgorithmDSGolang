package avl

type Node struct {
	val         int
	height      int
	left, right *Node
}

func NewRoot() *Node {
	return
}

func (n *Node) Get(val int) *Node {
	if n == nil {
		return nil
	}
	if n.val == val {
		return n
	}

	if val <= n.val {
		return n.left.Get(val)
	} else {
		return n.right.Get(val)
	}

}

func Insert(n *Node, val int) {
	if n == nil {
		n = &Node{
			val:    val,
			height: 0,
		}
		return
	}

	if val <= n.val {
		Insert(n.left, val)
	} else {
		Insert(n.right, val)
	}

	n.height = height(n)
	reBalance(n)
}

func reBalance(n *Node) {
	bFactor := balanceFactor(n)
	if bFactor == 2 { // Left heavy
		bFactor = balanceFactor(n.left)
		if bFactor == 1 { // LL rotation
			rlRotation(root)
		} else if bFactor == -1 { // LR rotation
			rrRotation(root)
		}
	} else if bFactor == -2 { //Right heavy
		bFactor = balanceFactor(n.right)
		if bFactor == 1 {
			rlRotation(root)
		} else if bFactor == -1 {
			rrRotation(root)
		}
	}
}

func Delete(n *Node, val int) {
	if n == nil {
		return
	}
	if n.val < val {
		Delete(n.right, val)
	} else if n.val > val {
		Delete(n.left, val)
	} else {
		// current node is to be deleted
		// cases:
		// 1. no child
		// 2. has one child
		// 3. has two child
		if n.left == nil && n.right == nil {
			n = nil
		} else if n.left == nil {
			n = n.right
		} else if n.right == nil {
			n = n.left
		} else {
			// shift bigger values down by one
			minVal := min(n.right)
			n.val = minVal
			Delete(n, minVal)
		}
	}

	n.height = height(n)
	reBalance(n)
}

func height(n *Node) int {
	if n == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 1
	}

	var leftHeight, rightHeight int
	if root.left != nil {
		leftHeight = height(left)
	}
	if root.right != nil {
		rightHeight = height(right)
	}

	max := leftHeight
	if rightHeight > leftHeight {
		max = rightHeight
	}

	return max + 1
}

func min(n *Node) int {
	if n == nil {
		return 0
	}

	if n.left == nil {
		return n.val
	}

	return min(n.left)
}

func balanceFactor(n *Node) int {
	return height(n.right) - height(n.left)
}
