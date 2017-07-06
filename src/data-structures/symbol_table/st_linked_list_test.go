package symbol_table

import (
	"testing"
)

var collection = map[string]int{
	"key1": 1,
	"key2": 2,
	"key3": 3,
}

func TestSTLinkedList(t *testing.T) {
	stll := NewSTLinkedList()
	for k, v := range collection {
		stll.Put(k, v)
	}

	for k, v := range collection {
		if stll.Get(k) != v {
			t.Errorf("STLinkedList fail to work")
		}
	}
}

func TestSTLinkedList_Put(t *testing.T) {
	stll := NewSTLinkedList()
	stll.Put(1, 1)
	stll.Put(1, 2)
	stll.Put(1, 3)
	if stll.Get(1) != 3 {
		t.Error("STLinkedList_Put error")
	}
}
