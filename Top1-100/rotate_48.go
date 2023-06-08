package main

// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[7,4,1],[8,5,2],[9,6,3]]

func rotate(matrix [][]int) {
	// 原地旋转，找规律，控制内外层遍历次数
	// 如果开新数组的话会非常方便，竖着从下向上输出即可（题目不允许）
	// 获取长度n
	length := len(matrix[0])
	// 得到最大下标，用来控制范围
	tmlen := length - 1
	// 内层变化
	aviaber := length
	// 外层限制，length小于3，外层便利1次，到下标0即可，其他的若为偶数，遍历次数为length/2-1,奇数为length/2
	var waviber int
	if length <= 3 {
		waviber = 0
	} else if length%2 == 0 {
		waviber = length/2 - 1
	} else {
		waviber = length / 2
	}
	// 双层循环遍历
	for i := 0; i <= waviber; i++ {
		// aviaber用来控制内层最大便利次数，每次减2，如果小于等于2 直接为0
		if aviaber > 2 {
			aviaber = aviaber - 2
		} else {
			aviaber = 0
		}
		// 内存便利最大下标为当前横坐标加上需要遍历的次数即可
		for j := i; j <= aviaber+i; j++ {
			// 用来原地对4个位置进行变换
			matrix[i][j], matrix[j][tmlen-i], matrix[tmlen-i][tmlen-j], matrix[tmlen-j][i] = matrix[tmlen-j][i], matrix[i][j], matrix[j][tmlen-i], matrix[tmlen-i][tmlen-j]
		}
	}
}
