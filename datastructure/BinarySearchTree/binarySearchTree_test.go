package binarySearchTree

import "testing"

func TestPrint(t *testing.T) {
	bst := NewBST()
	bst.Add(1)
	bst.Add(2)
	bst.Add(3)
	bst.Add(7)
	bst.Print()
}

func TestToStringInOrder(t *testing.T) {
	bst := NewBST()
	bst.Add(2)
	bst.Add(1)
	bst.Add(3)
	bst.Add(7)

	exp := " 1  2  3  7 "
	rst := bst.ToStringInOrder()
	if rst != exp {
		t.Errorf("result %s is different from expected %s", rst, exp)
	}
}

func TestToStringPostOrder(t *testing.T) {
	bst := NewBST()
	bst.Add(2)
	bst.Add(1)
	bst.Add(3)
	bst.Add(7)

	exp := " 1  7  3  2 "
	rst := bst.ToStringPostOrder()
	if rst != exp {
		t.Errorf("result %s is different from expected %s", rst, exp)
	}
}

func TestToStringPreOrder(t *testing.T) {
	bst := NewBST()
	bst.Add(2)
	bst.Add(1)
	bst.Add(3)
	bst.Add(7)

	exp := " 2  1  3  7 "
	rst := bst.ToStringPreOrder()
	if rst != exp {
		t.Errorf("result %s is different from expected %s", rst, exp)
	}
}

func TestToStringByLevel(t *testing.T) {
	bst := NewBST()
	bst.Add(2)
	bst.Add(1)
	bst.Add(3)
	bst.Add(7)

	bst.ToStringByLevel()
}
