package face17

// 面试题 17.04. 消失的数字
// 数组nums包含从0到n的所有整数，但其中缺了一个。请编写代码找出那个缺失的整数。你有办法在O(n)时间内完成吗？

// 示例 1：
// 输入：[3,0,1]
// 输出：2

// 示例 2：
// 输入：[9,6,4,2,3,5,7,0,1]
// 输出：8

// 使用异或运算，5^4^5 = 4 因此，从0~n 全部异或 再异或数组里的数据 即可得到缺失数据
func missingNumber(nums []int) int {
	length := len(nums)
	var sult int
	// 得到数组和对应范围的全部异或数据
	for i := 0; i < length; i++ {
		sult = sult ^ i ^ nums[i]
	}
	// 最后与最大值异或即可
	return sult ^ length
}
