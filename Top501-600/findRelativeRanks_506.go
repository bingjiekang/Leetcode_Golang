package top501600

import (
	"sort"
	"strconv"
)

// 506. 相对名次
// Type :简单

// 给你一个长度为 n 的整数数组 score ，其中 score[i] 是第 i 位运动员在比赛中的得分。所有得分都 互不相同 。
// 运动员将根据得分 决定名次 ，其中名次第 1 的运动员得分最高，名次第 2 的运动员得分第 2 高，依此类推。运动员的名次决定了他们的获奖情况：
// 名次第 1 的运动员获金牌 "Gold Medal" 。
// 名次第 2 的运动员获银牌 "Silver Medal" 。
// 名次第 3 的运动员获铜牌 "Bronze Medal" 。
// 从名次第 4 到第 n 的运动员，只能获得他们的名次编号（即，名次第 x 的运动员获得编号 "x"）。
// 使用长度为 n 的数组 answer 返回获奖，其中 answer[i] 是第 i 位运动员的获奖情况。
// 示例 1：
// 输入：score = [5,4,3,2,1]
// 输出：["Gold Medal","Silver Medal","Bronze Medal","4","5"]
// 解释：名次为 [1st, 2nd, 3rd, 4th, 5th] 。

// 排序数组，使用map存储对应分数对应的名次，最后再遍历原切片，对应数字输入切片数组即可
// 前三名 对应特殊名词
var target []string = []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) (sult []string) {
	// 存储分数对应的位数
	var dict map[int]string = make(map[int]string, 0)
	length := len(score)
	var nums []int = make([]int, length)
	// 深拷贝，socre复制给nums，注意他们的容量大小，左边大于等于右边，否则会少数据
	copy(nums, score)
	// 倒序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for k, v := range nums {
		if k < 3 {
			dict[v] = target[k]
		} else {
			dict[v] = strconv.FormatInt(int64(k+1), 10)
		}

	}
	for _, v := range score {
		sult = append(sult, dict[v])
	}
	return
}
