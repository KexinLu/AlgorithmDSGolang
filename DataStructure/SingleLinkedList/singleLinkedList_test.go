package SingleLinkedList

import (
	"testing"
)

func TestAddToHead(t *testing.T) {
	t.Run("Can Add To Linked List", func(t *testing.T) {
		ll := NewSingleLinkedList()
		ll.AddToHead(1)
		ll.AddToHead(2)
		ll.AddToHead(3)
		ll.AddToHead(4)

		expected := "4,3,2,1"
		rst := ll.ToString()

		if rst != expected {
			t.Errorf("Was expecting %s but got %s", expected, rst)
		}
	})
}

func TestAddToTail(t *testing.T) {
	t.Run("Can Add To Linked List Tail", func(t *testing.T) {
		ll := NewSingleLinkedList()
		ll.AddToTail(1)
		ll.AddToTail(2)
		ll.AddToTail(3)
		ll.AddToTail(4)

		expected := "1,2,3,4"
		rst := ll.ToString()

		if rst != expected {
			t.Errorf("Was expecting %s but got %s", expected, rst)
		}
	})
}

func TestRemoveHead(t *testing.T) {
	ll := NewSingleLinkedList()
	ll.AddToTail(1)
	ll.AddToTail(2)
	ll.AddToTail(3)
	ll.AddToTail(4)
	ll.RemoveHead()

	expected := "2,3,4"
	rst := ll.ToString()

	if rst != expected {
		t.Errorf("Was expecting %s but got %s", expected, rst)
	}
}

func TestRemoveTail(t *testing.T) {
	ll := NewSingleLinkedList()
	ll.AddToTail(1)
	ll.AddToTail(2)
	ll.AddToTail(3)
	ll.AddToTail(4)
	ll.RemoveTail()

	expected := "1,2,3"
	rst := ll.ToString()

	if rst != expected {
		t.Errorf("Was expecting %s but got %s", expected, rst)
	}
}

func TestReverse(t *testing.T) {
	ll := NewSingleLinkedList()
	ll.AddToTail(1)
	ll.AddToTail(2)
	ll.AddToTail(3)
	ll.AddToTail(4)
	ll.Reverse()

	expected := "4,3,2,1"
	rst := ll.ToString()

	if rst != expected {
		t.Errorf("Was expecting %s but got %s", expected, rst)
	}
}
