package main

// 给定一个方阵，其中每个单元(像素)非黑即白。设计一个算法，找出 4 条边皆为黑色像素的最大子方阵。
// 返回一个数组 [r, c, size] ，其中 r, c 分别代表子方阵左上角的行号和列号，size 是子方阵的边长。若有多个满足条件的子方阵，返回 r 最小的，若 r 相同，返回 c 最小的子方阵。若无满足条件的子方阵，返回空数组。

// 示例 1:
// 输入:
// [
//    [1,0,1],
//    [0,0,1],
//    [0,0,1]
// ]
// 输出: [1,0,2]
// 解释: 输入中 0 代表黑色，1 代表白色，标粗的元素即为满足条件的最大子方阵

// 解析：最大的黑方阵，从左上角遍历下标到右下角，判断不同长度的四条边是否都为0，如果都为0，且边的长度大于记录的长度，则将左上角下标及长度记录下来
func findSquare(matrix [][]int) (sult []int) {
	// 定义左上角为（x,y）然后对应的右上角（x,y+l），右下角(x+l,y+l)、左下角（x+l,y）l为1~length-1
	length := len(matrix)
	// vlx,vly,maxl分别用来记录左上角下标以及边的长度
	var vlx, vly, maxl int
	// x从0到length-1
	for x := 0; x < length; x++ {
		// y从0到length-1
		for y := 0; y < length; y++ {
			// l的长度从0到x|y +l<length
			for l := 0; y+l < length && x+l < length; l++ {
				// 用来判断这四个点之间，两两连线是否都为0
				if Seek(x, y, x, y+l, matrix) && Seek(x+l, y, x+l, y+l, matrix) && Seek(x, y+l, x+l, y+l, matrix) && Seek(x, y, x+l, y, matrix) {
					// 如果都为0，且边的长度大于maxl，则记录左上角下标以及长度
					if l+1 > maxl {
						maxl = l + 1
						vlx = x
						vly = y
					}
				}
			}
		}
	}
	// 如果maxl不为0，则证明存在都为0的边，则将其加入即可
	if maxl != 0 {
		sult = []int{vlx, vly, maxl}
	}
	// 对不存在都为0的边，直接返回空
	return

}

// 用来判断两点之间（即边）是否都为0
func Seek(x0, y0, x1, y1 int, lt [][]int) bool {
	if x0 == x1 {
		for i := y0; i <= y1; i++ {
			if lt[x0][i] != 0 {
				return false
			}
		}
		return true
	} else if y0 == y1 {
		for j := x0; j <= x1; j++ {
			if lt[j][y0] != 0 {
				return false
			}
		}
		return true
	}
	return false
}
