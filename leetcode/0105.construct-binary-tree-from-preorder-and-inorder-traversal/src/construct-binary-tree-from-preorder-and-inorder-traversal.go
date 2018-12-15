package src

// 根据一棵树的前序遍历与中序遍历构造二叉树

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	//前序遍历的第一个节点即根节点
	res := &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return res
	}

	// 在中序遍历中，根节点前面的即根节点的左子树，后面的即右子树

	//匿名函数
	idx := func(val int, nums []int) int {
		for i, v := range nums {
			if v == val {
				return i
			}
		}
		return -1
	}(res.Val, inorder)
	if idx == -1 {
		panic("data error")
	}

	//递归,构建根节点左子树和右子树
	res.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	res.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return res
}
