package main

import (
	"sort"
)

// 给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差
// 示例：

// 输入：{1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
// 输出：3，即数值对(11, 8)

// 1 <= a.length, b.length <= 100000
// -2147483648 <= a[i], b[i] <= 2147483647
// 正确结果在区间 [0, 2147483647] 内

// 解析先对两个数组进行排序,然后双指针进行移动比较,计算并记录差值,如果差值为0则直接返回,如果一个切片下标指向的数小,则其向后移动,反复比较,直至一个数到达末尾,因为是升序排列,所以逐个比较可行.
func smallestDifference(a []int, b []int) int {
	// lengtha/b 获取a,b切片的长度
	lengtha := len(a)
	lengthb := len(b)
	// 对a,b切片进行排序
	sort.Ints(a)
	sort.Ints(b)
	// dx记录两个值的差值
	var dx int = 2147483647
	// temp用来存储临时的值
	var temp int = 2147483647
	// i,j分别遍历a,b切片
	for i, j := 0, 0; i < lengtha && j < lengthb; {
		// 对比对应下标数据值,记录差值,将小的数后移
		if a[i] < b[j] {
			temp = b[j] - a[i]
			i++
		} else if a[i] > b[j] {
			temp = a[i] - b[j]
			j++
		} else { // 如果两个值差值为0直接返回即可
			return 0
		}
		// 如果差值大于得到的差值,记录最新小的值
		if dx >= temp {
			dx = temp
		}
	}
	return dx
}
