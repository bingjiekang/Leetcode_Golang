package main

// 给定二叉树的根节点 root ，返回所有左叶子之和。

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解析:通过深度遍历,直到遍历到左子叶然后加入到结果中,左子树遍历完后遍历右子树,直到遍历出所有的左子叶相加即可

// 主函数用来得到结果,根节点为空则直接返回
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

// dfs深度遍历,用来计算左子叶的和
func dfs(root *TreeNode) (sult int) {
	// 先遍历当前节点的左节点,如果不为空
	if root.Left != nil {
		// 则查看左节点的 左右子节点 如果都为空则证明这是叶子节点,加入即可
		if Leaf(root.Left) {
			sult += root.Left.Val
		} else { // 如果有一个不为空则继续遍历左节点,并将原来的结果保持并加入
			sult += dfs(root.Left)
		}
	}
	// 然后如果右节点不为空则遍历右节点
	if root.Right != nil {
		// 将结果保持并加入
		sult += dfs(root.Right)
	}
	return
}

// 用来查找当前节点的左右子节点是否存在,若都不存在则返回true,否则返回false
func Leaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
