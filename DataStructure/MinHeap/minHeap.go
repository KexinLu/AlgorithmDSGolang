package minHeap

type minHeap struct {
	heapArray []int
	size      int
	maxSize   int
}

func NewMinhp(maxSize int) *minHeap {
	return &minHeap{
		heapArray: []int{},
		size:      0,
		maxSize:   maxSize,
	}
}

// check if given index is a leaf node
// Note: in golang divide will auto floor
func (m *minHeap) isLeafIdx(idx int) bool {
	// [1,2,3] size 3, 3/2 = 1, everything has index greater than 0 should be leaf
	// thus if index >= size/2, it is a leaf node

	return idx >= m.size/2 && idx < m.size
}

func (m *minHeap) isNode(idx int) bool {
	return idx < m.size
}

// get given index's parent index
func (m *minHeap) parentIdx(idx int) int {
	if idx >= m.size || idx <= 2 {
		return 0
	}

	return (idx - 1) / 2
}

// get left child index
func (m *minHeap) leftChildIdx(idx int) int {
	return idx*2 + 1
}

func (m *minHeap) rightChildIdx(idx int) int {
	return idx*2 + 2
}

func (m *minHeap) swap(i, j int) {
	m.heapArray[i], m.heapArray[j] = m.heapArray[j], m.heapArray[i]
}

func (m *minHeap) upHeapify(idx int) {
	if idx >= m.size || idx == 0 {
		return
	}

	parentIdx := m.parentIdx(idx)
	if m.heapArray[idx] >= m.heapArray[parentIdx] {
		return
	} else {
		m.swap(idx, parentIdx)
		m.upHeapify(parentIdx)
	}
	return
}

func (m *minHeap) downHeapify(idx int) {
	if m.isLeafIdx(idx) || idx < 0 {
		return
	}

	minIdx := idx
	leftIdx := m.leftChildIdx(idx)
	rightIdx := m.rightChildIdx(idx)
	if m.isNode(leftIdx) && m.heapArray[minIdx] > m.heapArray[leftIdx] {
		minIdx = leftIdx
	}
	if m.isNode(rightIdx) && m.heapArray[minIdx] > m.heapArray[rightIdx] {
		minIdx = rightIdx
	}
	if minIdx == idx {
		return
	} else {
		m.swap(minIdx, idx)
		m.downHeapify(minIdx)
	}
	return
}

func (m *minHeap) insert(val int) {
	m.heapArray = append(m.heapArray, val)
	m.size++
	m.upHeapify(m.size - 1)
}

func (m *minHeap) remove() int {
	if m.size == 0 {
		return 0
	}

	top := m.heapArray[0]
	m.swap(0, m.size-1)
	m.heapArray = m.heapArray[:m.size-1]
	m.size--
	m.downHeapify(0)
	return top
}

func (m *minHeap) build() {
	for i := m.size/2 - 1; i >= 0; i-- {
		m.downHeapify(i)
	}
}
