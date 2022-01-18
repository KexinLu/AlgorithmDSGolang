package maxHeap

import (
	"fmt"
	"testing"
)

func TestCreatingMaxHeap(t *testing.T) {
	mh := CreateMaxHeap()
	data := []int{5, 2, 3, 1, 5}
	size := 5
	mh.SetData(data)
	mh.SetSize(size)

	if mh.size != 5 {
		t.Errorf("MaxHeap should have size 5 but got %v", mh.size)
	}
}

func TestFindParentIdx(t *testing.T) {
	mh := CreateMaxHeap()
	data := []int{5, 2, 3, 1, 5}
	size := 5
	mh.SetData(data)
	mh.SetSize(size)

	testCases := []struct {
		input  int
		expect int
	}{{4, 1}, {3, 1}, {2, 0}, {1, 0}}

	for _, tc := range testCases {
		if tc.expect != mh.parentIdx(tc.input) {
			t.Errorf("parent index %v is different from expected %v with input %v", mh.parentIdx(tc.input), tc.expect, tc.input)
		}
	}

}

func TestUpHeapify(t *testing.T) {
	mh := CreateMaxHeap()
	data := []int{5, 2, 3, 1, 6}
	size := 5
	mh.SetData(data)
	mh.SetSize(size)

	mh.upHeapify(4)
	expected := []int{6, 5, 3, 1, 2}
	if !arrayEquals(expected, mh.data) {
		t.Errorf("two arrays are different expected %v, actual %v", expected, mh.data)
	}
}

func TestSortList(t *testing.T) {
	sorted := SortSlice([]int{5, 6, 2, 1, 3, 4, 7})
	fmt.Println(sorted)
}

func arrayEquals(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for idx, item := range arr1 {
		if arr2[idx] != item {
			return false
		}
	}
	return true
}
