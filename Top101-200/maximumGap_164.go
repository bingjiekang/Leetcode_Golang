package main

import (
	"sort"
)

// 给定一个无序的数组 nums，返回 数组在排序之后，相邻元素之间最大的差值 。如果数组元素个数小于 2，则返回 0 。

// 示例 1:

// 输入: nums = [3,6,9,1]
// 输出: 3
// 解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。

// 通过排序然后对排序后的数组遍历,得到相邻的最大差值即为这个结果
// 实现sort的内部接口函数,就可以使用sort函数,使用sort函数需要实现Swap/Less/Len 三个函数

func maximumGap(nums []int) (sult int) {
	// 获取数组长度
	length := len(nums)
	// 长度小于2时直接返回0
	if length < 2 {
		return 0
	}
	// 定义一个结构体类型的temp
	var temp ST
	// temp里的数组进行赋值
	temp.nums = nums
	// 进行排序
	sort.Sort(sm(temp))
	// 得到最大间隙
	for i := 0; i < length-1; i++ {
		sult = max(sult, temp.nums[i+1]-temp.nums[i])
	}
	// 返回即可
	return
}

// 用来选择两者中最大的数
func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

// 定义一个整形数组的结构体
type ST struct {
	nums []int
}

// 定义一个ST类型的数据sm
type sm ST

// 对sm进行函数实现
// 对Len函数进行实现:返回待排序数组的长度
func (this sm) Len() int {
	return len(this.nums)
}

// 对Swap函数的实现:返回两不同数据的下标交换
func (this sm) Swap(n, m int) {
	this.nums[n], this.nums[m] = this.nums[m], this.nums[n]
}

// 实现降序和升序输出:小于时返回true为升序,大于时返回降序
func (this sm) Less(i, j int) bool {
	return this.nums[i] < this.nums[j]
}
