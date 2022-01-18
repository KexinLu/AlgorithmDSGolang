package minHeap

import (
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	mhp := NewMinhp(5)
	if mhp.maxSize != 5 {
		t.Errorf("minheap max size %v is different from expected %v ", mhp.maxSize, 5)
	}

	if mhp.size != 0 {
		t.Errorf("minheap size %v is different from expected %v ", mhp.size, 0)
	}
}

func TestIsLeaf(t *testing.T) {
	mhp := NewMinhp(5)
	mhp.heapArray = []int{1, 2, 3, 4, 5, 6}
	mhp.size = 6

	testCases := []struct {
		idx int
		exp bool
	}{
		{0, false},
		{1, false},
		{2, false},
		{3, true},
		{4, true},
		{5, true},
	}

	for _, tc := range testCases {
		if rst := mhp.isLeafIdx(tc.idx); rst != tc.exp {
			t.Errorf("node with index %v isLeafIdx result %v is different from expected %v", tc.idx, rst, tc.exp)
		}
	}
}

func TestGetParentIndex(t *testing.T) {
	mhp := NewMinhp(5)
	mhp.heapArray = []int{1, 2, 3, 4, 5, 6}
	mhp.size = 6

	testCases := []struct {
		idx int
		exp int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 1},
		{5, 2},
	}

	for _, tc := range testCases {
		if rst := mhp.parentIdx(tc.idx); rst != tc.exp {
			t.Errorf("node with index %v parentIdx result %v is different from expected %v", tc.idx, rst, tc.exp)
		}
	}
}

func TestSwap(t *testing.T) {
	mhp := NewMinhp(5)
	mhp.heapArray = []int{1, 2, 3, 4, 5, 6}
	mhp.size = 6

	mhp.swap(1, 3)
	expected := []int{1, 4, 3, 2, 5, 6}
	if !Equal(mhp.heapArray, expected) {
		t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
	}

}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestUpHeapify(t *testing.T) {
	mhp := NewMinhp(5)
	mhp.heapArray = []int{1, 2, 3, 4, 5, 6}
	mhp.size = 6

	t.Run("Can handle index out of bound", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5, 6}
		mhp.upHeapify(15)
		if !Equal(mhp.heapArray, expected) {
			t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
		}
	})

	t.Run("Can handle root index", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5, 6}
		mhp.upHeapify(0)
		if !Equal(mhp.heapArray, expected) {
			t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
		}
	})

	t.Run("Can bubble up index", func(t *testing.T) {
		mhp.heapArray = []int{1, 2, 3, 4, 5, 0}
		expected := []int{0, 2, 1, 4, 5, 3}
		mhp.upHeapify(5)
		if !Equal(mhp.heapArray, expected) {
			t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
		}
	})

	t.Run("Can bubble up non-leaf index", func(t *testing.T) {
		mhp.heapArray = []int{1, 0, 3, 2, 5, 4}
		expected := []int{0, 1, 3, 2, 5, 4}
		mhp.upHeapify(1)
		if !Equal(mhp.heapArray, expected) {
			t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
		}
	})
}

func TestDownHeapify(t *testing.T) {
	mhp := NewMinhp(5)
	mhp.heapArray = []int{6, 1, 2, 3, 4, 5}
	mhp.size = 6

	t.Run("Can handle index out of bound", func(t *testing.T) {
		expected := []int{6, 1, 2, 3, 4, 5}
		mhp.downHeapify(15)
		if !Equal(mhp.heapArray, expected) {
			t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
		}
	})

	t.Run("Can handle leaf index", func(t *testing.T) {
		expected := []int{6, 1, 2, 3, 4, 5}
		mhp.downHeapify(3)
		if !Equal(mhp.heapArray, expected) {
			t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
		}
	})

	t.Run("Can bubble down index", func(t *testing.T) {
		expected := []int{1, 3, 2, 6, 4, 5}
		mhp.downHeapify(0)
		if !Equal(mhp.heapArray, expected) {
			t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
		}
	})

	t.Run("Can bubble down non-leaf index", func(t *testing.T) {
		mhp.heapArray = []int{1, 6, 2, 3, 4, 5}
		expected := []int{1, 3, 2, 6, 4, 5}
		mhp.downHeapify(1)
		if !Equal(mhp.heapArray, expected) {
			t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
		}
	})

	t.Run("Can bubble down case 2", func(t *testing.T) {
		expected := []int{2, 4, 3, 5}
		mhp.size = 4
		mhp.heapArray = []int{5, 2, 3, 4}
		mhp.downHeapify(0)
		if !Equal(mhp.heapArray, expected) {
			t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
		}
	})
}

func TestInsert(t *testing.T) {
	mhp := NewMinhp(15)
	mhp.heapArray = []int{1, 2, 3, 4, 5}
	mhp.size = 5

	mhp.insert(0)
	expected := []int{0, 2, 1, 4, 5, 3}
	if !Equal(mhp.heapArray, expected) {
		t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
	}
}

func TestRemove(t *testing.T) {
	mhp := NewMinhp(15)
	mhp.heapArray = []int{1, 2, 3, 4, 5}
	mhp.size = 5

	expected := []int{2, 4, 3, 5}

	mhp.remove()

	if !Equal(mhp.heapArray, expected) {
		t.Errorf("heapArray %v is different from expected %v", mhp.heapArray, expected)
	}

}
