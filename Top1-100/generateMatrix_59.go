package top1100

// 59. 螺旋矩阵 II
// Type: 中等
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

// 示例 1：
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]

// 输入：n = 1
// 输出：[[1]]

// 对二维数组进行初始化，确定移动方向1-右，2-下，3-左,4-上，对螺旋排列的范围进行确定，确定边缘范围，以及到达边缘或者非0值后进行顺时针旋转继续移动，最后返回即可
func generateMatrix(n int) (sult [][]int) {
	sult = make([][]int, 0)
	// 对二维切片的数据进行初始化
	for i := 0; i < n; i++ {
		var lt []int = make([]int, n)
		sult = append(sult, lt)
	}
	// 1-右，2-下，3-左,4-上
	// 确定螺旋排列的最大范围，以及行列坐标
	maxSize := n * n
	row, col := 0, 0
	mv := 1
	// 对切片进行赋值
	for value := 1; value <= maxSize; value++ {
		sult[row][col] = value
		switch mv {
		case 1: // 1 则向右移动 col++
			col++
		case 2: // 2 向下移动 row++
			row++
		case 3: // 3 向左移动 col--
			col--
		case 4: // 4 向上移动 row--
			row--
		}
		// 如果到达边界或者非0值，则更新需要移动的下标，并且更改移动方向
		if col >= n || (mv == 1 && sult[row][col] != 0) {
			mv = 2
			col--
			row++
		} else if row >= n || (mv == 2 && sult[row][col] != 0) {
			mv = 3
			row--
			col--
		} else if col <= -1 || (mv == 3 && sult[row][col] != 0) {
			mv = 4
			col++
			row--
		} else if row <= -1 || (mv == 4 && sult[row][col] != 0) {
			mv = 1
			row++
			col++
		}
	}
	return
}
