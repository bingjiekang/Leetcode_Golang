package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。本题中，一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1
*/
/*思路
自底向上递归:先是从根节点往下,然后递归回溯的时候会比较每个根节点的左右子树的高度,若不满足则返回-1.类似于后序遍历，对于当前遍历到的节点，先递归地判断其左右子树是否平衡，再判断以当前节点为根的子树是否平衡。如果一棵子树是平衡的，则返回其高度（高度一定是非负整数），否则返回 -1。如果存在一棵子树不平衡,则整个二叉树一定不平衡。
*/

// 如果深度大于等于0，证明二叉树平衡
func isBalanced(root *TreeNode) bool {
	if Height(root) >= 0 {
		return true
	}
	return false
}

// 用来获取每个根节点左右子树的深度
func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lheight := Height(root.Left)
	rheight := Height(root.Right)
	// 如果深度不平衡则返回的会是-1，递归返回后lheight和rheight会受影响所以需要小于0
	if abs(lheight-rheight) > 1 || lheight < 0 || rheight < 0 {
		return -1
	}
	// 返回当前的根节点的最大深度
	return max(lheight, rheight) + 1
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return (-1) * num
}

func max(left int, right int) int {
	if left >= right {
		return left
	}
	return right
}
