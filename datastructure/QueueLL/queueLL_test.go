package queueLL

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(5)
	q.Enqueue(4)
	q.Enqueue(3)

	l := q.Length()
	if l != 3 {
		t.Errorf("queue length is %d which is different from expected %d", l, 3)
	}
}

func TestPeak(t *testing.T) {
	q := NewQueue()

	q.Enqueue(5)
	q.Enqueue(4)
	q.Enqueue(3)

	expected := 5
	rst, _ := q.Peak()

	if expected != rst {
		t.Errorf("node value %d is different from expected %d", rst, expected)
	}

	if q.Length() != 3 {
		t.Errorf("queue length should not have changed from %d to %d", 3, q.Length())
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(5)
	q.Enqueue(4)
	q.Enqueue(3)

	expected := 5
	rst, _ := q.Dequeue()

	if expected != rst {
		t.Errorf("node value %d is different from expected %d", rst, expected)
	}

	expected2 := 4
	rst2, _ := q.Peak()

	if expected2 != rst2 {
		t.Errorf("node value %d is different from expected %d", rst2, expected2)
	}

	if q.Length() != 2 {
		t.Errorf("queue length be %d but got %d", 2, q.Length())
	}
}
