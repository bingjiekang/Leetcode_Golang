package main

// 要求:给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历:根左右
func preorderTraversal(root *TreeNode) (list []int) {
	// 内部递归调用
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 根左右
		list = append(list, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return list
}
