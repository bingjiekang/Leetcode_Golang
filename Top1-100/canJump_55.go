package main

// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

// 示例 1：

// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

// 定义一个变量用来确定最远范围,遍历下标，如果出现下标加上其对应的值 大于当前最远范围，则更新，一直遍历，直到遍历到下标为far最远范围，或者是nums最后一个数组，如果下标遍历到最后一位，则返回true，否则返回false
func canJump(nums []int) bool {
	// far用来存储最远下标
	var far int = nums[0]
	length := len(nums) - 1
	// 遍历i，i小于最远下标且小于数组的长度
	for i := 0; i <= far && i <= length; i++ {
		// 如果遍历到最后一位则返回true
		if i == length {
			return true
		}
		// 如果当前可以到达的最远距离大于之前的距离，则更新最远距离
		if far < i+nums[i] {
			far = i + nums[i]
		}
	}
	return false
}
