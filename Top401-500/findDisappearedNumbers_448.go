package top401500

// 448. 找到所有数组中消失的数字
// 简单
// 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

// 示例 1：
// 输入：nums = [4,3,2,7,8,2,3,1]
// 输出：[5,6]
// 示例 2：
// 输入：nums = [1,1]
// 输出：[2]

func findDisappearedNumbers(nums []int) (st []int) {
	// 数组存储
	length := len(nums)
	var sult []int = make([]int, length)
	for _, v := range nums {
		sult[v-1] = 1
	}
	for k, v := range sult {
		if v == 0 {
			st = append(st, k+1)
		}
	}
	return
}
