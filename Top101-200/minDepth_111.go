package main

import (
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*二叉树的最小深度
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。
*/
/*深度搜索
首先可以想到使用深度优先搜索的方法，遍历整棵树，记录最小深度。
对于每一个非叶子节点，我们只需要分别计算其左右子树的最小叶子节点深度。这样就将一个大问题转化为了小问题，可以递归地解决该问题。
*/

func minDepth(root *TreeNode) int {
	// 没有根节点
	if root == nil {
		return 0
	}
	// 只有根节点
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// 根节点和左节点
	min_depth := math.MaxInt64
	if root.Left != nil {
		min_depth = min(minDepth(root.Left), min_depth)
	}
	// 根节点和右节点
	if root.Right != nil {
		min_depth = min(minDepth(root.Right), min_depth)
	}
	return min_depth + 1
}

func min(left int, right int) int {
	if left <= right {
		return left
	}
	return right
}
