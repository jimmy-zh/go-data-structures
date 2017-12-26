package symbol_table

type STLinkedList struct {
	root *node
	n    int
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
	stll.n++
}

func (stll *STLinkedList) Get(key interface{}) interface{} {
	for node := stll.root; node != nil; node = node.next {
		if node.key == key {
			return node.value
		}
	}
	return nil
}

func (stll *STLinkedList) Delete(key interface{}) {
	stll.root = stll.deleteRecurse(stll.root, key)
}

func (stll *STLinkedList) deleteRecurse(node *node, key interface{}) *node {
	if node == nil {
		return nil
	}
	if node.key == key {
		stll.n--
		return node.next
	}
	node.next = stll.deleteRecurse(node.next, key)
	return node
}

func (stll *STLinkedList) Empty() bool {
	return stll.n == 0
}

func (stll *STLinkedList) Size() int {
	return stll.n
}
