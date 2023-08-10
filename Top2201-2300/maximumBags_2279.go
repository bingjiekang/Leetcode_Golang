package main

import (
	"sort"
)

// 现有编号从 0 到 n - 1 的 n 个背包。给你两个下标从 0 开始的整数数组 capacity 和 rocks 。第 i 个背包最大可以装 capacity[i] 块石头，当前已经装了 rocks[i] 块石头。另给你一个整数 additionalRocks ，表示你可以放置的额外石头数量，石头可以往 任意 背包中放置。

// 请你将额外的石头放入一些背包中，并返回放置后装满石头的背包的 最大 数量。
// 示例 1：

// 输入：capacity = [2,3,4,5], rocks = [1,2,4,4], additionalRocks = 2
// 输出：3
// 解释：
// 1 块石头放入背包 0 ，1 块石头放入背包 1 。
// 每个背包中的石头总数是 [2,3,4,4] 。
// 背包 0 、背包 1 和 背包 2 都装满石头。
// 总计 3 个背包装满石头，所以返回 3 。
// 可以证明不存在超过 3 个背包装满石头的情况。
// 注意，可能存在其他放置石头的方案同样能够得到 3 这个结果。

// 解析:获得每个背包空的位置，然后排序，从小到大计数直到数量大于额外石头数量,返回满背包的个数

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	// 获取capcity的长度
	length := len(capacity)
	// 定义cout用来计数满背包的个数,sult用来记录当前空的个数是否超过最大待放入的个数
	var cout, sult int
	// 记录每个背包空的个数
	for i := 0; i < length; i++ {
		rocks[i] = capacity[i] - rocks[i]
	}
	// 排序
	sort.Ints(rocks)
	// 遍历这个切片,从前往后从小到大,计数背包的空间累计和
	for _, v := range rocks {
		sult += v
		// 如果背包空间小于等于待加入空间个数,计数满背包的个数
		if sult <= additionalRocks {
			cout++
		} else { // 否则直接退出
			break
		}
	}
	// 返回即可
	return cout
}
