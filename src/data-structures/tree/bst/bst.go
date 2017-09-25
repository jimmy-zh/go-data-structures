package bst

import "algorithms/searching/binary"

type Node struct {
	N           int //nodes in subtree rooted here
	Value       interface{}
	Key         binary.Comparabler
	left, right *Node
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

func (bst *BST) Size() int {
	return bst.root.Size()
}

func (bst *BST) Search(key binary.Comparabler) *Node {
	return bst.searchRecurse(bst.root, key)
}

func (bst *BST) searchRecurse(node *Node, key binary.Comparabler) *Node {
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

func (bst *BST) Insert(key binary.Comparabler, value interface{}) {
	bst.root = bst.insertRecurse(bst.root, key, value)
}

func (bst *BST) insertRecurse(node *Node, key binary.Comparabler, value interface{}) *Node {
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

func (bst *BST) Floor(key binary.Comparabler) *Node {
	return bst.floorRecurse(bst.root, key)
}

func (bst *BST) floorRecurse(node *Node, key binary.Comparabler) *Node {
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

func (bst *BST) Ceiling(key binary.Comparabler) *Node {
	return bst.ceilingRecurse(bst.root, key)
}

func (bst *BST) ceilingRecurse(node *Node, key binary.Comparabler) *Node {
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
