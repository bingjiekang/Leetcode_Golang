package main

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

// 问总共有多少条不同的路径？

// 输入：m = 3, n = 7
// 输出：28

// 动态规划:m行n列的话,第0行和第0列都为1,其他的值 都为左边一格加上边一格得到的值即lt[i][j-1]+lt[i-1][j],最后返回lt[m-1][n-1]即可
// func uniquePaths(m int, n int) int {
// 	if m == 1 || n == 1 {
// 		return 1
// 	}

// 	var lt [][]int = make([][]int, 0)
// 	for i := 0; i < m; i++ {
// 		lt = append(lt, []int{})
// 		for j := 0; j < n; j++ {
// 			if i == 0 || j == 0 {
// 				lt[i] = append(lt[i], 1)
// 			} else {
// 				lt[i] = append(lt[i], lt[i][j-1]+lt[i-1][j])
// 			}
// 		}
// 	}
// 	return lt[m-1][n-1]
// }

// 压缩为一维段数组,新的每一行得到的数据,更新上一次得到的数据
func uniquePaths(m int, n int) int {
	// 如果为1,则直接返回1
	if m == 1 || n == 1 {
		return 1
	}
	// 否则就遍历一定的长度
	m, n = m-1, n-1
	var lt []int = make([]int, 0)
	st := m * n
	// yt用来记录从理论上第二行开始,第一列的数
	var yt int = 3
	// 遍历从0到st-1
	for i := 0; i < st; i++ {
		// 如果是理论上的第一行,则直接更新对应值
		if i < n {
			lt = append(lt, i+2)
		} else if i%n == 0 { // 如果是理论上的其他行,则更新对应第一列的值
			lt[i%n] = yt
			yt++
		} else { // 理论上的其他行非第一列的值,更新覆盖原数组为当前下标值%n和当前下标%n-1的值
			lt[i%n] = lt[i%n-1] + lt[i%n]
		}
	}
	return lt[n-1]
}
