package src

// TreeNode Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//4种情况：
	if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
		// 1. p,q分别在root的左右两边，那么 乘积小于0
		//2. p或者q中一方==root，另一方为root的子节点（孙节点。。。）
		return root
	} else if root.Val > p.Val {
		//3. p，q都在root的左子树中
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		//4. p，q都在root的右子树中
		return lowestCommonAncestor(root.Right, p, q)
	}

}
