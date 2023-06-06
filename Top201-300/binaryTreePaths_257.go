package main

import (
	"strconv"
)

// 你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
// 叶子节点 是指没有子节点的节点。

// 输入：root = [1,2,3,null,5]
// 输出：["1->2->5","1->3"]
// 解析:最直观的方法是使用深度优先搜索。在深度优先搜索遍历二叉树时，我们需要考虑当前的节点以及它的孩子节点。
// 如果当前节点不是叶子节点，则在当前的路径末尾添加该节点，并继续递归遍历该节点的每一个孩子节点。
// 如果当前节点是叶子节点，则在当前路径末尾添加该节点后我们就得到了一条从根节点到叶子节点的路径，将该路径加入到答案即可。
// 如此，当遍历完整棵二叉树以后我们就得到了所有从根节点到叶子节点的路径

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义全局字符串列表，存储不同的路径
var tm []string

func binaryTreePaths(root *TreeNode) []string {
	// 每次调用的时候重新更新数据表为空列表
	tm = []string{}
	DfsPath(root, "")
	return tm
}

func DfsPath(root *TreeNode, path string) {
	// 如果当前节点不为空
	if root != nil {
		// 用来存储当前路径下当前节点
		path = path + strconv.Itoa(root.Val)
		// 如果左子叶和右子叶都为空，则证明这是当前路径的最后一个元素，则将当前节点加入到列表中
		if root.Left == nil && root.Right == nil {
			tm = append(tm, path)
			return
		} else { // 如果不为空则继续遍历，深度遍历
			path += "->"
			DfsPath(root.Left, path)  // 先遍历左子树，
			DfsPath(root.Right, path) // 再遍历右子树
		}
	}
}
