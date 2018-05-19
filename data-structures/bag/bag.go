package bag

import "sync"

type Bag struct {
	m     sync.RWMutex
	items []interface{}
}

func NewBag() *Bag {
	return &Bag{
		items: make([]interface{}, 0),
	}
}

func (b *Bag) Add(v interface{}) {
	b.m.Lock()
	defer b.m.Unlock()
	b.items = append(b.items, v)
}

func (b *Bag) Size() int {
	b.m.RLock()
	defer b.m.RUnlock()
	return len(b.items)
}

func (b *Bag) Empty() bool {
	return b.Size() == 0
}

func (b *Bag) Items() []interface{} {
	b.m.RLock()
	defer b.m.RUnlock()
	return b.items
}
