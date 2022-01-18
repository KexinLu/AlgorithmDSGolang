package BTree

import "testing"

func TestCreateTree(t *testing.T) {
	bt := CreateBTree()
	if bt.size != 0 {
		t.Error("binary tree should have size 0 ")
	}
}

func TestInsertToEmptyTree(t *testing.T) {
	bt := CreateBTree()
	bt.Insert(15)
	if bt.size != 1 {
		t.Error("binary tree should have size 1 ")
	}

	if bt.root.value != 15 {
		t.Errorf("root should have value %v but got %v", 15, bt.root.value)
	}
}
