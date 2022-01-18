package maxHeap

import "fmt"

type MaxHeap struct {
	data []int
	size int
}

func CreateMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: []int{},
	}
}

func (mh *MaxHeap) forceVal(idx int, val int) {
	mh.data[idx] = val
}

func (mh *MaxHeap) SetData(data []int) {
	mh.data = data
}

func (mh *MaxHeap) SetSize(size int) {
	mh.size = size
}

func (mh *MaxHeap) isLeaf(idx int) bool {
	return idx > mh.size/2
}

func (mh *MaxHeap) parentIdx(idx int) int {
	return (idx - 1) / 2
}

func (mh *MaxHeap) leftIdx(idx int) int {
	return idx*2 + 1
}

func (mh *MaxHeap) rightIdx(idx int) int {
	return idx*2 + 2
}

func (mh *MaxHeap) upHeapify(idx int) {
	if idx == 0 {
		return
	}
	parentIdx := mh.parentIdx(idx)
	if mh.data[parentIdx] > mh.data[idx] {
		return
	}

	// bubbling up
	mh.data[parentIdx], mh.data[idx] = mh.data[idx], mh.data[parentIdx]
	mh.upHeapify(parentIdx)
}

func (mh *MaxHeap) downHeapify(idx int) {
	if mh.isLeaf(idx) || idx >= mh.size-1 {
		return
	}
	leftIdx := mh.leftIdx(idx)
	rightIdx := mh.rightIdx(idx)
	max := idx

	if leftIdx < mh.size && mh.data[idx] < mh.data[leftIdx] {
		max = leftIdx
	}
	if rightIdx < mh.size && mh.data[max] < mh.data[rightIdx] {
		max = rightIdx
	}

	if max == idx {
		return
	}

	mh.data[max], mh.data[idx] = mh.data[idx], mh.data[max]
	mh.downHeapify(idx)
}

func (mh *MaxHeap) buildHeap() {
	for cur := (mh.size - 1) >> 1; cur >= 0; cur-- {
		mh.downHeapify(cur)
	}
}

func (mh *MaxHeap) Swap(a, b int) {
	mh.data[a], mh.data[b] = mh.data[b], mh.data[a]
}

func SortSlice(input []int) []int {
	mh := CreateMaxHeap()
	size := len(input)
	mh.SetSize(size)
	mh.SetData(input)
	mh.buildHeap()
	for last := len(input) - 1; last > 0; last-- {
		mh.Swap(0, last)
		size--
		mh.SetSize(size)
		mh.buildHeap()
		fmt.Println(mh.data)
	}

	return mh.data
}
