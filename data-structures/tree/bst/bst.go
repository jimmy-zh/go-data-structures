package bst

import (
	"github.com/midnight-vivian/go-data-structures/utils"
)

type Node struct {
	Key   utils.Comparabler
	Value interface{}
	left  *Node
	right *Node
	N     int //nodes in subtree rooted here
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return n.N
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return new(BST)
}

func (bst *BST) Search(key utils.Comparabler) *Node {
	return bst.searchRecurse(bst.root, key)
}

func (bst *BST) searchRecurse(node *Node, key utils.Comparabler) *Node {
	if node == nil {
		return nil
	}
	switch {
	case key.More(node.Key):
		return bst.searchRecurse(node.right, key)
	case key.Less(node.Key):
		return bst.searchRecurse(node.left, key)
	default:
		return node
	}
}

func (bst *BST) Insert(key utils.Comparabler, value interface{}) {
	bst.root = bst.insertRecurse(bst.root, key, value)
}

func (bst *BST) insertRecurse(node *Node, key utils.Comparabler, value interface{}) *Node {
	if node == nil {
		node = &Node{Key: key, Value: value, N: 1}
	} else {
		switch {
		case key.More(node.Key):
			node.right = bst.insertRecurse(node.right, key, value)
		case key.Less(node.Key):
			node.left = bst.insertRecurse(node.left, key, value)
		default:
			node.Value = value
		}
		node.N = node.left.Size() + node.right.Size() + 1
	}
	return node
}

func (bst *BST) Remove(key utils.Comparabler) {
	bst.root = bst.removeRecurse(bst.root, key)
}

func (bst *BST) removeRecurse(node *Node, key utils.Comparabler) *Node {
	if node == nil {
		return nil
	}
	switch {
	case key.More(node.Key):
		node.right = bst.removeRecurse(node.right, key)
	case key.Less(node.Key):
		node.left = bst.removeRecurse(node.left, key)
	default:
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		t := node
		node = bst.minRecurse(t.right)
		node.right = bst.removeMinRecurse(t.right)
		node.left = t.left
	}
	node.N = node.left.Size() + node.right.Size() + 1
	return node
}

func (bst *BST) Size() int {
	return bst.root.Size()
}

func (bst *BST) Max() *Node {
	if bst.Size() == 0 {
		return nil
	}
	return bst.maxRecurse(bst.root)
}

func (bst *BST) maxRecurse(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return bst.maxRecurse(node.right)
}

func (bst *BST) Min() *Node {
	if bst.Size() == 0 {
		return nil
	}
	return bst.minRecurse(bst.root)
}

func (bst *BST) minRecurse(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return bst.minRecurse(node.left)
}

func (bst *BST) RemoveMax() {
	bst.root = bst.removeMaxRecurse(bst.root)
}

func (bst *BST) removeMaxRecurse(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node.left
	}
	node.right = bst.removeMaxRecurse(node.right)
	node.N = node.left.Size() + node.right.Size() + 1
	return node
}

func (bst *BST) RemoveMin() {
	bst.root = bst.removeMinRecurse(bst.root)
}

func (bst *BST) removeMinRecurse(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node.right
	}
	node.left = bst.removeMinRecurse(node.left)
	node.N = node.left.Size() + node.right.Size() + 1
	return node
}

func (bst *BST) Floor(key utils.Comparabler) *Node {
	return bst.floorRecurse(bst.root, key)
}

func (bst *BST) floorRecurse(node *Node, key utils.Comparabler) *Node {
	if node == nil {
		return nil
	}
	switch {
	case key.More(node.Key):
		if n := bst.floorRecurse(node.right, key); n == nil {
			return node
		} else {
			return n
		}
	case key.Less(node.Key):
		return bst.floorRecurse(node.left, key)
	default:
		return node
	}
}

func (bst *BST) Ceiling(key utils.Comparabler) *Node {
	return bst.ceilingRecurse(bst.root, key)
}

func (bst *BST) ceilingRecurse(node *Node, key utils.Comparabler) *Node {
	if node == nil {
		return nil
	}
	switch {
	case key.Less(node.Key):
		if n := bst.ceilingRecurse(node.left, key); n == nil {
			return node
		} else {
			return n
		}
	case key.More(node.Key):
		return bst.ceilingRecurse(node.right, key)
	default:
		return node
	}
}

func (bst *BST) Rank(key utils.Comparabler) int {
	return bst.rankRecurse(bst.root, key)
}

func (bst *BST) rankRecurse(node *Node, key utils.Comparabler) int {
	if node == nil {
		return 0
	}
	switch {
	case key.More(node.Key):
		return bst.rankRecurse(node.right, key) + node.left.Size() + 1
	case key.Less(node.Key):
		return bst.rankRecurse(node.left, key)
	default:
		return node.left.Size()
	}
}

func (bst *BST) Select(k int) *Node {
	if k < 0 || k >= bst.Size() {
		return nil
	}
	return bst.selectRecurse(bst.root, k)
}

func (bst *BST) selectRecurse(node *Node, k int) *Node {
	if node == nil {
		return nil
	}
	s := node.left.Size()
	switch {
	case k < s:
		return bst.selectRecurse(node.left, k)
	case k > s:
		return bst.selectRecurse(node.right, k-s-1)
	default:
		return node
	}
}
