package main

// 给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

// 输入：nums = [9,6,4,2,3,5,7,0,1]
// 输出：8
// 解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。

// 方法:用数学方法求和,没有出现的,用总体和的结果减去数组中的内容就是最后消失的数字
func missingNumber(nums []int) int {
	// 获取数组的长度
	length := len(nums)
	// 计算前n个数的总和
	sn := length * (1 + length) / 2
	// 找到缺失的数
	for _, v := range nums {
		sn -= v
	}
	return sn
}
