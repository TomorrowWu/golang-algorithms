### 描述
根据一棵树的前序遍历与中序遍历构造二叉树

**注意:**

你可以假设树中没有重复的元素。

例如，给出

```
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
```

返回如下的二叉树：
```
    3
   / \
  9  20
    /  \
   15   7
```

### 思路
- 前序遍历中，第一个节点即根节点
- 在中序遍历中，找出第一个节点的位置，根节点前面的 L 个数据，即根节点左子树的中序遍历数据，前序遍历中根节点后面的 L 个数据即左子树的前序遍历
- 右子树同上
- 简而言之，确定了根节点，确定了左子树和右子树的数据，递归对左子树和右子树进行重建

### 代码实现
```Golang
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
```

### 题目来源
[105. 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

### GitHub
- [源码传送门](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0105.construct-binary-tree-from-preorder-and-inorder-traversal/src/construct-binary-tree-from-preorder-and-inorder-traversal.go)
- 各种数据结构及算法实现, LeetCode解题思路及答案
- [大厂面试题汇总及答案解析](https://github.com/TomorrowWu/interview)

> 本文为原创文章，转载注明出处，欢迎扫码关注公众号 ```楼兰``` 或者网站[https://lovecoding.club](https://lovecoding.club),第一时间看后续精彩文章，觉得好的话，顺手分享到朋友圈吧，感谢支持。

![image](https://upload-images.jianshu.io/upload_images/5815624-4a8b49cfbaf037dd.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/200)
