package main

// 给定一个整数数组，找出总和最大的连续数列，并返回总和。

// 示例：

// 输入： [-2,1,-3,4,-1,2,1,-5,4]
// 输出： 6
// 解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。

// 利用动态规划，总和最大的连续数列，假设前i段的和为f(i)，比较f(i)+nums[i] 和nums[i]的大小，如果f(i)+nums[i]大，则更新当前f(i)为f(i)+nums[i] 比较这个值是否大于sult，如果大于则更新，否则不更新。如果nums[i]大，则更新f(i)为nums[i]，通过从前往后遍历即可。
// 分析：因为想要总和最大的连续数列，那么这个序列的两端需要为正（除非都是负数，找最小的），前面一段和本段序列的第一位比较后，如果前面的加上本位的大于本位的，则证明前面的可以用，如果前面的加上本位的不大于本位，就从当前开始，动态规划可以解决这个问题
func maxSubArray(nums []int) int {
	// 获取长度
	length := len(nums)
	// ever为前面的f(i)，now为nums[i]
	var ever, now int
	// sult为最后的结果 总和最大数列
	var sult int = nums[0]
	// 从第一个开始到最后
	for i := 0; i < length; i++ {
		now = nums[i]
		// 更新ever为f(i)+nums[i]
		ever = ever + nums[i]
		// 如果f(i)+nums[i]>nums[i]
		if ever > now { // ever不变，便于后续累加，当前值更改
			now = ever
		} else { // 如果 f(i)+nums[i]<=nums[i]
			ever = now // 更新ever为当前值，从当前为再开始
		}
		// 当前值大于sult则更新
		if now > sult {
			sult = now
		}
	}
	return sult
}
