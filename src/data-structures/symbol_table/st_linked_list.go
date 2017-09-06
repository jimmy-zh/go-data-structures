package symbol_table

type STLinkedList struct {
	root *node
}

type node struct {
	next  *node
	key   interface{}
	value interface{}
}

func NewSTLinkedList() *STLinkedList {
	return &STLinkedList{}
}

func (stll *STLinkedList) Put(key, value interface{}) {
	for node := stll.root; node != nil; node = node.next {
		if node.key == key {
			node.value = value
			return
		}
	}
	stll.root = &node{
		next:  stll.root,
		key:   key,
		value: value,
	}
}

func (stll *STLinkedList) Get(key interface{}) interface{} {
	for node := stll.root; node != nil; node = node.next {
		if node.key == key {
			return node.value
		}
	}
	return nil
}
