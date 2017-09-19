package symbol_table

import (
	"algorithms/searching/binary"
)

type STTwoArrays struct {
	keys   []binary.Comparabler
	values []interface{}
}

func NewSTTwoArrays() *STTwoArrays {
	return &STTwoArrays{
		keys:   make([]binary.Comparabler, 0),
		values: make([]interface{}, 0),
	}
}

func (stta *STTwoArrays) rank(key binary.Comparabler) int {
	return binary.Search(stta.keys, key)
}

func (stta *STTwoArrays) Get(key binary.Comparabler) interface{} {
	index := stta.rank(key)
	if index < len(stta.keys) && stta.keys[index] == key {
		return stta.values[index]
	}
	return nil
}

func (stta *STTwoArrays) Put(key binary.Comparabler, value interface{}) {
	index := stta.rank(key)
	if index < len(stta.keys) && stta.keys[index] == key {
		stta.values[index] = value
		return
	}
	stta.keys = append(append(stta.keys[:index], key), stta.keys[index:]...)
	stta.values = append(append(stta.values[:index], value), stta.values[index:]...)
}
