package queueLL

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type QueueLL struct {
	head   *Node
	tail   *Node
	length int
}

func NewQueue() *QueueLL {
	return &QueueLL{}
}

func (q *QueueLL) Enqueue(val int) {
	newNode := &Node{
		value: val,
	}

	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		q.length++

		return
	}

	newNode.prev = q.tail
	q.tail.next = newNode
	q.tail = newNode
	q.length++
}

func (q *QueueLL) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}

	rstNode := q.head
	if rstNode.next != nil {
		rstNode.next.prev = nil
	}

	q.head = rstNode.next
	q.length--
	return rstNode.value, true
}

func (q *QueueLL) IsEmpty() bool {
	return q.length == 0
}

func (q *QueueLL) Peak() (int, bool) {
	if q.head == nil {
		return 0, false
	}

	return q.head.value, true
}

func (q *QueueLL) Length() int {
	return q.length
}
