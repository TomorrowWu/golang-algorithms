package src

//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced 是否平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalanced(root.Left) && isBalanced(root.Right) && abs(maxDepth(root.Left)-maxDepth(root.Right)) <= 1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
