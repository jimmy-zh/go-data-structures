package stack

import "sync"

type StackSlice struct {
	m     sync.RWMutex
	items []interface{}
}

func NewStackSlice() *StackSlice {
	return &StackSlice{
		items: make([]interface{}, 0),
	}
}

func (ss *StackSlice) Push(v interface{}) {
	ss.m.Lock()
	defer ss.m.Unlock()
	ss.items = append(ss.items, v)
}

func (ss *StackSlice) Pop() interface{} {
	if ss.Empty() {
		return nil
	}
	ss.m.Lock()
	defer ss.m.Unlock()
	l := len(ss.items)
	lastIn := ss.items[l-1]
	ss.items = ss.items[:l-1]
	return lastIn
}

func (ss *StackSlice) Peek() interface{} {
	if ss.Empty() {
		return nil
	}
	ss.m.RLock()
	defer ss.m.RUnlock()
	return ss.items[len(ss.items)-1]
}

func (ss *StackSlice) Size() int {
	ss.m.RLock()
	defer ss.m.RUnlock()
	return len(ss.items)
}

func (ss *StackSlice) Empty() bool {
	ss.m.RLock()
	ss.m.RUnlock()
	return len(ss.items) == 0
}
