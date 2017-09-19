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

func TestNewSTTwoArrays(t *testing.T) {
	tests := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	stta := NewSTTwoArrays()
	if stta.Get(String("key1")) != nil {
		t.Error("STTwoArrays get error")
	}
	for k, v := range tests {
		stta.Put(String(k), v)
	}
	for k, v := range tests {
		if stta.Get(String(k)).(string) != v {
			t.Error("STTwoArrays put error")
		}
	}

}

func TestSTTwoArrays_Put(t *testing.T) {
	stta := NewSTTwoArrays()
	stta.Put(String("key1"), "value1")
	stta.Put(String("key1"), "value2")

	if stta.Get(String("key1")) != "value2" {
		t.Error("STTwoArrays get error")
	}
}


