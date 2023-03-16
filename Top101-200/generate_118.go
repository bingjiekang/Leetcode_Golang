package main

/*杨辉三角
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和
*/
// 方法:分析当前要填的数据跟其他数据的关系特性
func generate(numRows int) [][]int {
	// 切片在使用前必须先声明，几纬声明几次
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
