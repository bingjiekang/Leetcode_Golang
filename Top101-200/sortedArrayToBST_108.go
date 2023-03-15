package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 将有序数组转换为二叉搜索树
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
*/
/*思路
选择中间(中间偏左)数字作为二叉搜索树的根节点,这样分给左右子树的数字个数相同或只相差1,可以使得树保持平衡。(left+right)/2
确定平衡二叉搜索树的根节点之后，其余的数字分别位于平衡二叉搜索树的左子树和右子树中，左子树和右子树分别也是平衡二叉搜索树，因此可以通过递归的方式创建平衡二叉搜索树。
*/

func sortedArrayToBST(nums []int) *TreeNode {
	// 定义递归函数，用来遍历左子树和右子树
	var dfs func(left int, right int) *TreeNode
	dfs = func(left int, right int) *TreeNode {
		// 越界退出
		if left > right {
			return nil
		}
		// 获取每个子节点的根节点，每个叶子节点都是对应的根节点
		mid := (left + right) / 2
		// 对每个根节点进行赋值
		temple := &TreeNode{
			Val:   nums[mid],
			Left:  dfs(left, mid-1),
			Right: dfs(mid+1, right),
		}
		// 返回当前叶子根节点
		return temple
	}
	return dfs(0, len(nums)-1)
}
