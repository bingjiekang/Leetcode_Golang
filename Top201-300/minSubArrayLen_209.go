package main

// 给定一个含有 n 个正整数的数组和一个正整数 target 。

// 找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
// 示例 1：

// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 示例 2：

// 输入：target = 4, nums = [1,4,4]
// 输出：1
// 示例 3：

// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0

// 使用滑动窗口
// 解析：使用两个指针确定左右两端的范围，右边为达到条件的最小范围，左边开始进行范围收缩，直到收缩到最短长度记录并直到有小于这个长度范围出现的时候再更新
func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	// sult为连续子数组的和，best为最短长度
	var sult, best int = 0, 1e5
	// 左右指针
	left, right := 0, 0
	// 右边范围小于length
	for ; right < length; right++ {
		// 更新连续子数组和的结果
		sult += nums[right]
		// 如果sult大于目标值，则开始移动左指针
		if sult >= target {
			// 移动左指针如果当前仍然满足条件则减去右边指针对应的值，再进行更新
			for left < right && sult-nums[left] >= target {
				sult -= nums[left]
				left++
			}
			// 如果此时的长度小于最短长度，则更新最短长度
			if right-left+1 < best {
				best = right - left + 1
			}
		}
	}
	// 如果最后sult满足大于等于target的条件 则返回最短长度
	if sult >= target {
		return best
	}
	// 否则返回0
	return 0
}
