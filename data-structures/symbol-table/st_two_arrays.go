package symbol_table

import (
	"github.com/midnight-vivian/go-data-structures/algorithms/searching/binary"
	"github.com/midnight-vivian/go-data-structures/utils"
)

var _ SymbolTableSuper = (*STTwoArrays)(nil)

type STTwoArrays struct {
	keys   []utils.Comparabler
	values []interface{}
	n      int
}

func NewSTTwoArrays() *STTwoArrays {
	return &STTwoArrays{
		keys:   make([]utils.Comparabler, 0),
		values: make([]interface{}, 0),
	}
}

func (stta *STTwoArrays) Get(key utils.Comparabler) interface{} {
	index := stta.Rank(key)
	if index < stta.n && stta.keys[index] == key {
		return stta.values[index]
	}
	return nil
}
func (stta *STTwoArrays) Put(key utils.Comparabler, value interface{}) {
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

func (stta *STTwoArrays) Delete(key utils.Comparabler) {
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

func (stta *STTwoArrays) Contains(key utils.Comparabler) bool {
	if stta.Get(key) != nil {
		return true
	}
	return false
}

func (stta *STTwoArrays) Max() utils.Comparabler {
	if stta.Empty() {
		return nil
	}
	return stta.keys[stta.n-1]
}

func (stta *STTwoArrays) Min() utils.Comparabler {
	if stta.Empty() {
		return nil
	}
	return stta.keys[0]
}

func (stta *STTwoArrays) DeleteMax() {
	if stta.n != 0 {
		stta.keys = stta.keys[:stta.n-1]
		stta.values = stta.values[:stta.n-1]
		stta.n--
	}
}

func (stta *STTwoArrays) DeleteMin() {
	if stta.n != 0 {
		stta.keys = stta.keys[1:]
		stta.values = stta.values[1:]
		stta.n--
	}
}

func (stta *STTwoArrays) Floor(key utils.Comparabler) utils.Comparabler {
	index := stta.Rank(key)
	if index < stta.n && stta.keys[index] == key {
		return stta.keys[index]
	}
	if index == 0 {
		return nil
	}
	return stta.keys[index-1]
}

func (stta *STTwoArrays) Ceiling(key utils.Comparabler) utils.Comparabler {
	index := stta.Rank(key)
	if index == stta.n {
		return nil
	}
	return stta.keys[index]
}

func (stta *STTwoArrays) Rank(key utils.Comparabler) int {
	if stta.Empty() {
		return 0
	}
	return binary.Search(stta.keys, key)
}

func (stta *STTwoArrays) Select(k int) utils.Comparabler {
	if k < stta.n {
		return stta.keys[k]
	}
	return nil
}
