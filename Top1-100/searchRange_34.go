package main

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

// 如果数组中不存在目标值 target，返回 [-1, -1]。

// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
// 示例 1：

// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
// 示例 2：

// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]

// 解析通过二分法:查找到第一次出现和最后一次出现的位置,如果没有出现返回-1,使用begin和end进行标记,如果在其中一个值没有改变,证明查询的值全部大于这个数组或者全部小于这个数组

func searchRange(nums []int, target int) []int {
	length := len(nums)
	left, right := 0, length-1
	// mid记录中间下标,begin和end记录第一次出现和最后一次出现的下标,初始值为-2用来控制当插入的数大于全部数组或者小于全部数组时,用来区分
	var mid, begin, end int = 0, -2, -2
	// 找到第一次出现目标数的下标-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
			begin = right
		}
	}
	mid = 0
	left, right = 0, length-1
	// 找到最后一次出现目标数的下标-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
			end = left
		} else {
			right = mid - 1
		}
	}
	// 用来判断其中一个是否没有变化,即为这个数大于或者小于全部数组
	if begin == -2 || end == -2 {
		return []int{-1, -1}
	} // 如果范围差值大于1,这证明有此下标,返回即可
	if end-begin > 1 {
		return []int{begin + 1, end - 1}
	}
	// 否则就是没有下标,返回-1
	return []int{-1, -1}

}
