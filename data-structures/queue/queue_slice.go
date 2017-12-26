package queue

import "sync"

type QueueSlice struct {
	m     sync.RWMutex
	items []interface{}
}

func NewQueueSlice() *QueueSlice {
	return &QueueSlice{
		items: make([]interface{}, 0),
	}
}

func (q *QueueSlice) Enqueue(v interface{}) {
	q.m.Lock()
	defer q.m.Unlock()
	q.items = append(q.items, v)
}

func (q *QueueSlice) Dequeue() interface{} {
	q.m.Lock()
	defer q.m.Unlock()
	v := q.items[0]
	q.items = q.items[1:]
	return v
}

func (q *QueueSlice) Size() int {
	q.m.RLock()
	defer q.m.RUnlock()
	return len(q.items)
}

func (q *QueueSlice) Empty() bool {
	q.m.RLock()
	defer q.m.RUnlock()
	return len(q.items) == 0
}
