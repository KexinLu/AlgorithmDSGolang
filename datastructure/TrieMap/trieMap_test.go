package trie

import (
	"testing"
)

func TestNewTrie(t *testing.T) {
	nt := NewTrie()

	if nt == nil {
		t.Error("trie is nil")
	}

	if nt.IsLeaf() {
		t.Error("root is not a leaf")
	}
}

func TestInsert(t *testing.T) {
	nt := NewTrie()
	nt.Insert("abc")
	nt.Insert("acd")
	nt.PrintByLayer()
}
