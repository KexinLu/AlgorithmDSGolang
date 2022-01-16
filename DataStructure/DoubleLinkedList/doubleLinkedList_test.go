package DoubleLinkedList

import (
	"fmt"
	"testing"
)

func TestAddToHead(t *testing.T) {
	dbl := NewDouble()

	dbl.AddToHead(1)
	dbl.AddToHead(2)
	dbl.AddToHead(3)
	dbl.AddToHead(4)
	fmt.Println(dbl.ToString())

	expected := "4,3,2,1"

	if rst := dbl.ToString(); rst != expected {
		t.Errorf("Was expecting %s but got %s", expected, rst)
	}
}

func TestAddToTail(t *testing.T) {
	dbl := NewDouble()

	dbl.AddToTail(1)
	dbl.AddToTail(2)
	dbl.AddToTail(3)
	dbl.AddToTail(4)

	expected := "1,2,3,4"

	if rst := dbl.ToString(); rst != expected {
		t.Errorf("Was expecting %s but got %s", expected, rst)
	}
}

func TestDeleteHead(t *testing.T) {
	dbl := NewDouble()

	dbl.AddToTail(1)
	dbl.AddToTail(2)
	dbl.AddToTail(3)
	dbl.AddToTail(4)
	dbl.DeleteHead()

	expected := "2,3,4"

	if rst := dbl.ToString(); rst != expected {
		t.Errorf("Was expecting %s but got %s", expected, rst)
	}
}

func TestDeleteTail(t *testing.T) {
	dbl := NewDouble()

	dbl.AddToTail(1)
	dbl.AddToTail(2)
	dbl.AddToTail(3)
	dbl.AddToTail(4)
	dbl.DeleteTail()

	expected := "1,2,3"

	if rst := dbl.ToString(); rst != expected {
		t.Errorf("Was expecting %s but got %s", expected, rst)
	}
}
