package main

// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
// 输入：n = 3
// 输出：5
// 示例 2：

// 输入：n = 1
// 输出：1
// 1<=n<=19
// 题目要求是计算不同二叉搜索树的个数。为此，我们可以定义两个函数：
// G(n): 长度为 n 的序列能构成的不同二叉搜索树的个数。
// F(i,n): 以 i 为根、序列长度为 n 的不同二叉搜索树个数 (1≤i≤n)。
// 可见，G(n)是我们求解需要的函数。

// G(n) = i从1到n 求和 F(i,n)
// G(0)=G(1) = 1
// 给定序列 1⋯n 我们选择数字 i 作为根，则根为 i的所有二叉搜索树的集合是左子树集合和右子树集合的笛卡尔积，对于笛卡尔积中的每个元素，加上根节点之后形成完整的二叉搜索树
// F(i,n)=G(i−1)⋅G(n−i)
// G(n)= i从1到n 求和 G(i−1)⋅G(n−i)

func numTrees(n int) int {
	// 用来存储对应位数的排列个数
	var G []int = make([]int, n+1)
	// 初始化 G[0]为1
	G[0] = 1
	// i从1到n循环
	for i := 1; i <= n; i++ {
		// 内层用来得到第i位对应的排列个数
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}
