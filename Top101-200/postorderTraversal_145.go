package main

// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历:左右根
func postorderTraversal(root *TreeNode) (list []int) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		list = append(list, root.Val)
	}
	dfs(root)
	return list
}
