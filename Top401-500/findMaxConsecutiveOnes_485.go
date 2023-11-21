package top401500

// 485. 最大连续 1 的个数
// Type :简单
// 给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。

// 示例 1：
// 输入：nums = [1,1,0,1,1,1]
// 输出：3
// 解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
// 示例 2:
// 输入：nums = [1,0,1,1,0,1]
// 输出：2
// 分别计数，如果出现0就判断是否为最大值，然后重新计算，最后由于最后一位为1则不会再进入，则最后再判断一次count是否大于max
func findMaxConsecutiveOnes(nums []int) int {
	var max, count int
	// 计数连续1出现的次数
	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	// 最后再更新一次得到结果即可
	if count > max {
		return count
	} else {
		return max
	}
}
