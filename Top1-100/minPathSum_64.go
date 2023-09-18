package main

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

// 说明：每次只能向下或者向右移动一步。
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
// 示例 2：

// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12

// 使用动态规划，由于只能向下或者向右，则更新第一行为当前值加上左边的值，更新第一列为当前值加上上面的值
// 非第一行或者第一列的，则更新当前值为min(左边，上边)加上当前值，遍历结束，最后一个值即为结果
func minPathSum(grid [][]int) int {
	// m,n为对应行和列的长度
	m, n := len(grid), len(grid[0])
	// 赋值temp变量接收grid
	temp := grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 { // i==0 && j==0 则不更新
				continue
			} else if i == 0 && j != 0 { // 如果i为0，j不为0更新当前值为 当前值加上左边的值
				temp[i][j] = temp[i][j] + temp[i][j-1]
			} else if j == 0 && i != 0 {
				temp[i][j] = temp[i][j] + temp[i-1][j] // 如果j为0，i不为0 更新当前值为 当前值加上上面的值
			} else {
				temp[i][j] = temp[i][j] + min(temp[i][j-1], temp[i-1][j]) // i,j都不为0，则更新当前值为min(左边，上边)加上当前值
			}
		}
	}
	// 返回结果即可
	return temp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
