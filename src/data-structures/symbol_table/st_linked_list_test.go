package symbol_table

import (
	"testing"
)

func TestSTLinkedList(t *testing.T) {
	tests := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	stll := NewSTLinkedList()
	for k, v := range tests {
		stll.Put(k, v)
	}

	for k, v := range tests {
		if stll.Get(k) != v {
			t.Errorf("STLinkedList fail to work")
		}
	}
}

func TestSTLinkedList_Put(t *testing.T) {
	stll := NewSTLinkedList()
	stll.Put(1, 1)
	stll.Put(1, 2)
	if stll.Get(1) != 2 {
		t.Error("STLinkedList_Put error")
	}
}
