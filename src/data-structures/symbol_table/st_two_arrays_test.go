package symbol_table

import (
	"testing"
)

type Key []string

func (k *Key) Get(i int) interface{} {
	return (*k)[i]
}

func (k *Key) Set(i int, v interface{}) {
	if i < k.Len() {
		*k = append(append((*k)[:i], v.(string)), (*k)[i:]...)
	} else {
		*k = append((*k)[:i], v.(string))
	}
}

func (k *Key) Len() int {
	return len(*k)
}

func (k *Key) Less(i int, t interface{}) bool {
	return (*k)[i] < t.(string)
}

func (k *Key) More(i int, t interface{}) bool {
	return (*k)[i] > t.(string)
}

func TestSTTwoArrays(t *testing.T) {
	tests := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	key := make(Key, 0)
	stta := NewSTTwoArrays(&key)
	if stta.Get("key1") != nil {
		t.Error("STTwoArrays get error")
	}
	for k, v := range tests {
		stta.Put(k, v)
	}
	for k, v := range tests {
		if stta.Get(k).(string) != v {
			t.Error("STTwoArrays set error")
		}
	}
}

func TestSTTwoArrays_Put(t *testing.T) {
	key := make(Key, 0)
	stta := NewSTTwoArrays(&key)

	stta.Put("key1", "value1")
	stta.Put("key1", "value2")

	if stta.Get("key1").(string) != "value2" {
		t.Error("STTwoArrays_Put error")
	}

}
