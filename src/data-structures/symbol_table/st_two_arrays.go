package symbol_table

import (
	"algorithms/searching/binary"
)

type Keyer interface {
	Get(i int) interface{}
	Set(i int, key interface{})
	binary.Comparabler
}

type STTwoArrays struct {
	keys   Keyer
	values []interface{}
}

func NewSTTwoArrays(keys Keyer) *STTwoArrays {
	return &STTwoArrays{
		keys:   keys,
		values: make([]interface{}, 0),
	}
}

func (stta *STTwoArrays) Get(key interface{}) interface{} {
	index := stta.rank(key)
	if index < stta.keys.Len() && stta.keys.Get(index) == key {
		return stta.values[index]
	}
	return nil
}

func (stta *STTwoArrays) Put(key, value interface{}) {
	index := stta.rank(key)
	if index < stta.keys.Len() {
		if stta.keys.Get(index) == key {
			stta.values[index] = value
		} else {
			stta.keys.Set(index, key)
			stta.values = append(append(stta.values[:index], value), stta.values[index:]...)
		}
	} else {
		stta.keys.Set(index, key)
		stta.values = append(stta.values[:index], value)
	}
}

func (stta *STTwoArrays) rank(key interface{}) int {
	return binary.Search(stta.keys, key)
}
