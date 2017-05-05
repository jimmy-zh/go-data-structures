package stack

type StackLinkedList struct {
	head *Node
	size int
}

type Node struct {
	next  *Node
	value interface{}
}

func NewStackLinkedList() *StackLinkedList {
	return new(StackLinkedList)
}

func (sll *StackLinkedList) Push(v interface{}) {
	headTmp := sll.head
	sll.head = &Node{
		next:  headTmp,
		value: v,
	}
	sll.size++
}

func (sll *StackLinkedList) Pop() interface{} {
	if sll.Empty() {
		return nil
	}
	headTmp := sll.head
	sll.head = sll.head.next
	sll.size--
	return headTmp.value
}

func (sll *StackLinkedList) Peek() interface{} {
	if sll.Empty() {
		return nil
	}
	return sll.head.value
}

func (sll *StackLinkedList) Size() int {
	return sll.size
}

func (sll *StackLinkedList) Empty() bool {
	return sll.size == 0
}
