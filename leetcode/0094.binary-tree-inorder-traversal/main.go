package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	if root != nil {
		ret = append(ret, inorderTraversal(root.Left)...)
		ret = append(ret, root.Val)
		ret = append(ret, inorderTraversal(root.Right)...)
	}
	return ret
}
