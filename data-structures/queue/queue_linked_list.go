package queue

type QueueLinkedList struct {
	first, last *Node
	size        int
}

type Node struct {
	next  *Node
	value interface{}
}

func NewQueueLinkedList() *QueueLinkedList {
	return new(QueueLinkedList)
}

func (q *QueueLinkedList) Enqueue(v interface{}) {
	node := &Node{
		value: v,
	}
	if q.Empty() {
		q.first, q.last = node, node
	} else {
		q.last.next = node
		q.last = node
	}
	q.size++
}

func (q *QueueLinkedList) Dequeue() interface{} {
	if q.Empty() {
		return nil
	}
	node := q.first
	q.first = q.first.next
	q.size--
	return node.value
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) Empty() bool {
	return q.Size() == 0
}
