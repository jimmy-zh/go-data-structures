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
	if value == nil {
		return
	}
	stb.bst.Insert(key, value)
}

func (stb *STBST) Delete(key binary.Comparabler) {

}

func (stb *STBST) Size() int {
	return stb.bst.Size()
}

func (stb *STBST) Empty() bool {
	return stb.Size() == 0
}

func (stb *STBST) Contains(key binary.Comparabler) bool {
	return stb.Get(key) != nil
}

func (stb *STBST) Max() binary.Comparabler {
	if n := stb.bst.Max(); n != nil {
		return n.Key
	}
	return nil
}

func (stb *STBST) Min() binary.Comparabler {
	if n := stb.bst.Min(); n != nil {
		return n.Key
	}
	return nil
}

func (stb *STBST) Floor(key binary.Comparabler) binary.Comparabler {
	if n := stb.bst.Floor(key); n != nil {
		return n.Key
	}
	return nil
}

func (stb *STBST) Ceiling(key binary.Comparabler) binary.Comparabler {
	if n := stb.bst.Ceiling(key); n != nil {
		return n.Key
	}
	return nil
}

//func (stb *STBST) Rank(key binary.Comparabler) int {

//}

func (stb *STBST) Select(k int) binary.Comparabler {
	if n := stb.bst.Select(k); n != nil {
		return n.Key
	}
	return nil
}
