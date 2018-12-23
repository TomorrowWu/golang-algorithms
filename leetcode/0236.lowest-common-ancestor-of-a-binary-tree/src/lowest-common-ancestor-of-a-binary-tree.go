package src

//TreeNode Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		//未找到p和q
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		//找到p或q
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		//p,q不存在
		return nil
	} else if left != nil && right != nil {
		//1， p,q分别在root的两边
		return root
	} else {
		//2, p,q在左子树或右子树
		if left != nil {
			//都在左子树
			return left
		}
		//都在右子树
		return right
	}
}
