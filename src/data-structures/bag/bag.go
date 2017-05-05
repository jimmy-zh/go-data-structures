package bag

type Bag struct {
	items []interface{}
	size  int
}

func NewBag() *Bag {
	return &Bag{
		items: make([]interface{}, 0),
	}
}

func (b *Bag) Add(v interface{}) {
	b.items = append(b.items, v)
}

func (b *Bag) Size() int {
	return b.size
}

func (b *Bag) Empty() bool {
	return b.Size() == 0
}
