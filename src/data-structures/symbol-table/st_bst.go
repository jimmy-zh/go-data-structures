package symbol_table

import (
	"github.com/midnight-vivian/go-data-structures/src/utils"

	"github.com/midnight-vivian/go-data-structures/src/data-structures/tree/bst"
)

var _ SymbolTableSuper = (*STBST)(nil)

type STBST struct {
	bst *bst.BST
}

func NewSTBST() *STBST {
	return &STBST{bst: bst.NewBST()}
}

func (stb *STBST) Get(key utils.Comparabler) interface{} {
	node := stb.bst.Search(key)
	if node == nil {
		return nil
	}
	return node.Value
}

func (stb *STBST) Put(key utils.Comparabler, value interface{}) {
	if value == nil {
		return
	}
	stb.bst.Insert(key, value)
}

func (stb *STBST) Delete(key utils.Comparabler) {
	stb.bst.Remove(key)
}

func (stb *STBST) Size() int {
	return stb.bst.Size()
}

func (stb *STBST) Empty() bool {
	return stb.Size() == 0
}

func (stb *STBST) Contains(key utils.Comparabler) bool {
	return stb.Get(key) != nil
}

func (stb *STBST) Max() utils.Comparabler {
	if n := stb.bst.Max(); n != nil {
		return n.Key
	}
	return nil
}

func (stb *STBST) Min() utils.Comparabler {
	if n := stb.bst.Min(); n != nil {
		return n.Key
	}
	return nil
}

func (stb *STBST) DeleteMax() {
	stb.bst.RemoveMax()
}

func (stb *STBST) DeleteMin() {
	stb.bst.RemoveMin()
}

func (stb *STBST) Floor(key utils.Comparabler) utils.Comparabler {
	if n := stb.bst.Floor(key); n != nil {
		return n.Key
	}
	return nil
}

func (stb *STBST) Ceiling(key utils.Comparabler) utils.Comparabler {
	if n := stb.bst.Ceiling(key); n != nil {
		return n.Key
	}
	return nil
}

func (stb *STBST) Rank(key utils.Comparabler) int {
	return stb.bst.Rank(key)
}

func (stb *STBST) Select(k int) utils.Comparabler {
	if n := stb.bst.Select(k); n != nil {
		return n.Key
	}
	return nil
}
