package main

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// 使用双向指针
func removeElement(nums []int, val int) int {
	length := len(nums)
	for left, right := 0, length-1; left <= right; {
		// 找到等于val的值
		for left < right && nums[left] != val {
			left++
		}
		// 找到不等于val的值
		for left <= right && nums[right] == val {
			right--
		}
		// 如果right在left左边或者和left相等，返回right+1即可
		if right <= left {
			return right + 1
		}
		// 否则的话就更新左边对应下标left的值
		nums[left] = nums[right]
		// 对right进行更新
		right--
	}
	return 0
}
