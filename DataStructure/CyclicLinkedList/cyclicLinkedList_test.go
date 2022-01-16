package cyclicLinkedList

import (
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	t.Run("Can create New LinkedList", func(t *testing.T) {
		l := NewLinkedList()
		expected := ""
		rst := l.ToString()

		if rst != expected {
			t.Errorf("was expecting %s but got %s", expected, rst)
		}
	})
}

func TestAddToHead(t *testing.T) {
	l := NewLinkedList()
	l.AddToHead(1)
	l.AddToHead(2)
	l.AddToHead(3)
	l.AddToHead(4)
	l.AddToHead(5)

	expected := "5,4,3,2,1"
	rst := l.ToString()

	if rst != expected {
		t.Errorf("was expecting %s but got %s", expected, rst)
	}
}

func TestAddToTail(t *testing.T) {
	l := NewLinkedList()
	l.AddToTail(1)
	l.AddToTail(2)
	l.AddToTail(3)
	l.AddToTail(4)
	l.AddToTail(5)

	expected := "1,2,3,4,5"
	rst := l.ToString()

	if rst != expected {
		t.Errorf("was expecting %s but got %s", expected, rst)
	}
}

func TestReverse(t *testing.T) {
	l := NewLinkedList()
	l.AddToTail(1)
	l.AddToTail(2)
	l.AddToTail(3)
	l.AddToTail(4)
	l.AddToTail(5)
	l.Reverse()

	expected := "5,4,3,2,1"
	rst := l.ToString()

	if rst != expected {
		t.Errorf("was expecting %s but got %s", expected, rst)
	}
}
