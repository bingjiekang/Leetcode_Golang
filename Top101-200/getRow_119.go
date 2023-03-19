package main

// 要求:给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
func getRow(rowIndex int) []int {
	// 将前n行数据生成
	lt := generate(rowIndex + 1)
	// 返回即可
	return lt[rowIndex]
}

// 用来生成前n行的杨辉三角二维数据表
func generate(numRows int) [][]int {
	// 切片在使用前必须先声明，几位声明几次
	list := make([][]int, numRows)
	// 横坐标从0开始到某一行
	for i := 0; i < numRows; i++ {
		// 内部切片声明
		list[i] = make([]int, i+1)
		// 对每个边缘进行赋值为1
		list[i][0], list[i][i] = 1, 1
		// 对中间的数根据杨辉三角的特性进行赋值
		for j := 1; j <= (i - 1); j++ {
			list[i][j] = list[i-1][j-1] + list[i-1][j]
		}
	}
	return list
}
