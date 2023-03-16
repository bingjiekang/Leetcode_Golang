package main

/* 计算路径总和
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false .
叶子节点 是指没有子节点的节点。
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 方法
询问是否存在从当前节点 root 到叶子节点的路径，满足其路径和为 sum。假定从根节点到当前节点的值之和为 val，我们可以将这个大问题转化为一个小问题：是否存在从当前节点的子节点到叶子的路径，满足其路径和为 sum - val。
*/
// 深度优先搜索，递归计算，从根节点到叶子节点的数之和等于目标的数，则根节点的左右节点的数到叶子节点的值应该为目标数减当前根节点数，可以递归调用
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 如果是空节点返回false
	if root == nil {
		return false
	}
	// 计算需要当前节点到叶子结点需要满足的数据值
	sum := targetSum - root.Val
	// sum=0则找到匹配的值但是同时需要该节点是叶子节点无左右子节点
	if sum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	// 左右子树只要有一个满足即可
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)

}
