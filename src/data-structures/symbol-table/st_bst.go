package symbol_table

import (
	"data-structures/tree/bst"
	"algorithms/searching/binary"
)

type STBST struct {
	bst *bst.BST
}

func NewSTBST() *STBST {
	return &STBST{bst: bst.NewBST()}
}

func (stb *STBST) Get(key binary.Comparabler) interface{} {
	node := stb.bst.Search(key)
	if node == nil {
		return nil
	}
	return node.Value
}

func (stb *STBST) Put(key binary.Comparabler, value interface{}) {
	stb.bst.Insert(key, value)
}

func (stb *STBST) Delete(key binary.Comparabler) {

}

func (stb *STBST) Max() binary.Comparabler {
	return stb.bst.Max().Key
}

func (stb *STBST) Min() binary.Comparabler {
	return stb.bst.Min().Key
}

func (stb *STBST) Floor(key binary.Comparabler) binary.Comparabler {

}
