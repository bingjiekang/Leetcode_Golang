package main

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 请必须使用时间复杂度为 O(log n) 的算法

// 示例 1:

// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
// 解析:通过二分查找,如果查找不到就证明需要插入的在mid右边或者mid本身
func searchInsert(nums []int, target int) int {
	// 二分查找
	length := len(nums)
	left, right := 0, length-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	// 如果nums[mid]小于target证明需要插入到mid+1
	if nums[mid] < target {
		return mid + 1
	} else { // 如果nums[mid]大于target证明需要插入到本位
		return mid
	}
}
