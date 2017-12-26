package utils

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreeConstructor(vals []interface{}) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, 0)
	parent := 0
	left := true
	node := (*TreeNode)(nil)
	for i := range vals {
		if val, ok := vals[i].(int); ok {
			node = &TreeNode{Val: val}
		} else {
			node = nil
		}
		nodes = append(nodes, node)
		if i == 0 {
			continue
		}
		for parent < i && nodes[parent] == nil {
			parent++
		}
		if parent == i {
			//invalid input vals
			return nil
		}
		if left {
			nodes[parent].Left = node
			left = false
		} else {
			nodes[parent].Right = node
			left = true
			parent++
		}
	}
	/*
		for _, node := range nodes {
			if node == nil {
				fmt.Println("null")
			} else {
				fmt.Println(node.Val)
			}
		}
	*/
	return nodes[0]
}

func BinaryTreeEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	if !BinaryTreeEqual(t1.Left, t2.Left) || !BinaryTreeEqual(t1.Right, t2.Right) {
		return false
	}
	return true
}
