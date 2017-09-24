package symbol_table

import (
	"testing"
)

type String string

func (s String) Less(item interface{}) bool {
	return s < item.(String)
}

func (s String) More(item interface{}) bool {
	return s > item.(String)
}

func TestSTTwoArrays_Get(t *testing.T) {
	tests := map[string]interface{}{
		"key1": "value1",
		"key2": 2,
		"key3": '3',
	}

	stta := NewSTTwoArrays()
	if stta.Get(String("key1")) != nil {
		t.Error("STTwoArrays get error")
	}
	for k, v := range tests {
		stta.Put(String(k), v)
	}
	for k, v := range tests {
		if stta.Get(String(k)) != v {
			t.Error("STTwoArrays get error")
		}
	}
}

func TestSTTwoArrays_Put(t *testing.T) {
	stta := NewSTTwoArrays()
	stta.Put(String("key1"), "value1")
	stta.Put(String("key1"), 2)

	if stta.Get(String("key1")) != 2 {
		t.Error("STTwoArrays get error")
	}
}

func TestSTTwoArrays_Delete(t *testing.T) {
	stta := NewSTTwoArrays()
	stta.Put(String("key1"), "value1")
	stta.Delete(String("key1"))
	if stta.Get(String("key1")) != nil {
		t.Error("STTwoArrays delete error")
	}
}

func TestSTTwoArrays_Mxx(t *testing.T) {
	stta := NewSTTwoArrays()
	if stta.Max() != nil {
		t.Error("STTwoArrays max error")
	}
	if stta.Min() != nil {
		t.Error("STTwoArrays min error")
	}

	stta.Put(String("key1"), "value1")
	stta.Put(String("key2"), "value2")

	if stta.Max() != String("key2") {
		t.Error("STTwoArrays max error")
	}
	if stta.Min() != String("key1") {
		t.Error("STTwoArrays min error")
	}
}
