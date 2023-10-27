package top19012000

// 1909. 删除一个元素使数组严格递增
// Type: 简单
// 给你一个下标从 0 开始的整数数组 nums ，如果 恰好 删除 一个 元素后，数组 严格递增 ，那么请你返回 true ，否则返回 false 。如果数组本身已经是严格递增的，请你也返回 true 。

// 数组 nums 是 严格递增 的定义为：对于任意下标的 1 <= i < nums.length 都满足 nums[i - 1] < nums[i] 。

// 示例 1：
// 输入：nums = [1,2,10,5,7]
// 输出：true
// 解释：从 nums 中删除下标 2 处的 10 ，得到 [1,2,5,7] 。
// [1,2,5,7] 是严格递增的，所以返回 true 。

// 遍历检查，如果存在非递增的就对这种现象加1，如果大于1就返回false，判断造成非递增的这两个数，
// nums[x-1],nums[x] 如果nums[x-1]>=nums[x+1] && nums[x]<=nums[x-2] 这种情况下 是无法严格递增的直接返回
// 其他依次遍历即可
func canBeIncreasing(nums []int) bool {
	length := len(nums)
	// 记录非递增的次数
	var record int
	for i := 1; i < length; i++ {
		// 非递减
		if nums[i-1] >= nums[i] {
			record++
			if record > 1 {
				return false
			}
			// 此情况下，直接返回false，因为删除i-1和删除i都无法使数组满足严格递增
			if i-2 >= 0 && i+1 < length && nums[i-1] >= nums[i+1] && nums[i] <= nums[i-2] {
				return false
			}
		}
	}
	return true
}
