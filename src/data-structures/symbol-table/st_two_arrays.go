package symbol_table

import (
	"algorithms/searching/binary"
)

type STTwoArrays struct {
	keys   []binary.Comparabler
	values []interface{}
	n      int
}

func NewSTTwoArrays() *STTwoArrays {
	return &STTwoArrays{
		keys:   make([]binary.Comparabler, 0),
		values: make([]interface{}, 0),
	}
}

func (stta *STTwoArrays) Get(key binary.Comparabler) interface{} {
	index := stta.Rank(key)
	if index < stta.n && stta.keys[index] == key {
		return stta.values[index]
	}
	return nil
}
func (stta *STTwoArrays) Put(key binary.Comparabler, value interface{}) {
	if value == nil {
		return
	}
	index := stta.Rank(key)
	if index < stta.n && stta.keys[index] == key {
		stta.values[index] = value
		return
	}
	stta.keys = append(append(stta.keys[:index], key), stta.keys[index:]...)
	stta.values = append(append(stta.values[:index], value), stta.values[index:]...)
	stta.n++
}

func (stta *STTwoArrays) Delete(key binary.Comparabler) {
	index := stta.Rank(key)
	if index < stta.n && stta.keys[index] == key {
		stta.keys = append(stta.keys[:index], stta.keys[index+1:]...)
		stta.values = append(stta.values[:index], stta.values[index+1:]...)
		stta.n--
	}
}

func (stta *STTwoArrays) Size() int {
	return stta.n
}

func (stta *STTwoArrays) Empty() bool {
	return stta.n == 0
}

func (stta *STTwoArrays) Contains(key binary.Comparabler) bool {
	if stta.Get(key) != nil {
		return true
	}
	return false
}

func (stta *STTwoArrays) Max() binary.Comparabler {
	if stta.Empty() {
		return nil
	}
	return stta.keys[stta.n-1]
}

func (stta *STTwoArrays) Min() binary.Comparabler {
	if stta.Empty() {
		return nil
	}
	return stta.keys[0]
}

func (stta *STTwoArrays) Floor(key binary.Comparabler) binary.Comparabler {
	index := stta.Rank(key)
	if index < stta.n && stta.keys[index] == key {
		return stta.keys[index]
	}
	if index == 0 {
		return nil
	}
	return stta.keys[index-1]
}

func (stta *STTwoArrays) Ceiling(key binary.Comparabler) binary.Comparabler {
	index := stta.Rank(key)
	if index == stta.n {
		return nil
	}
	return stta.keys[index]
}

func (stta *STTwoArrays) Rank(key binary.Comparabler) int {
	if stta.Empty() {
		return 0
	}
	return binary.Search(stta.keys, key)
}

func (stta *STTwoArrays) Select(k int) binary.Comparabler {
	if k < stta.n {
		return stta.keys[k]
	}
	return nil
}
