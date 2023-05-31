package main

// 要求:给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
// 左右节点翻转
// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]

// 解析:翻转节点,可以递归到子节点,然后对应翻转即可,然后逐步向上依次翻转从根节点开始，递归地对树进行遍历，并从叶子节点先开始翻转。如果当前遍历到的节点 root和root 的左右两棵子树都已经翻转，那么我们只需要交换两棵子树的位置，即可完成以 root 为根节点的整棵子树的翻转。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// 如果当前节点为空则返回
	if root == nil {
		return root
	}
	// 用来递归遍历左子树的子叶节点
	left := invertTree(root.Left)
	// 递归遍历右子树的子叶节点
	right := invertTree(root.Right)
	// 从底往上 一个个左右交换
	root.Left, root.Right = right, left
	// 返回交换后的根子节点
	return root
}
