package main

import (
	"strconv"
)

// 给定一个  无重复元素 的 有序 整数数组 nums 。
// 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
// 列表中的每个区间范围 [a,b] 应该按如下格式输出：
// "a->b" ，如果 a != b
// "a" ，如果 a == b

// 输入：nums = [0,1,2,4,5,7]
// 输出：["0->2","4->5","7"]
// 解释：区间范围是：
// [0,2] --> "0->2"
// [4,5] --> "4->5"
// [7,7] --> "7"

// 解析:一次便利,遇到正常增1的向后走,不满足的记录数据

func summaryRanges(nums []int) (result []string) {
	// 获取长度为0直接返回
	length := len(nums)
	if length == 0 {
		return
	}
	// 用来记录左右的开始和结束节点
	start, end := nums[0], 0
	for i := 1; i < length; i++ {
		// 如果当前值和前一个值之差不为1
		if nums[i]-nums[i-1] != 1 {
			// 前一个end值
			end = nums[i-1]
			// 如果相同,则表面是一个
			if start == end {
				result = append(result, strconv.Itoa(start))
			} else { // 不同则记录之前的start->end
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			// 更新start值为当前新开始的值
			start = nums[i]
		} else { // 为1的话,更新end为当前值
			end = nums[i]
		}
	}
	// 用来解决最后一位是单独的情况
	if start == nums[length-1] {
		result = append(result, strconv.Itoa(start))
	} else { // 解决后几位是连续的,由于前面已经记录了每一个start,则更新最后一个end即可
		result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}
	return
}
