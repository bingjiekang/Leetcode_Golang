package lcp

// LCP 66. 最小展台数量 简单
// 力扣嘉年华将举办一系列展览活动，后勤部将负责为每场展览提供所需要的展台。 已知后勤部得到了一份需求清单，记录了近期展览所需要的展台类型， demand[i][j] 表示第 i 天展览时第 j 个展台的类型。 在满足每一天展台需求的基础上，请返回后勤部需要准备的 最小 展台数量。
// 注意：同一展台在不同天中可以重复使用。

// 示例 1：
// 输入：demand = ["acd","bed","accd"]

// 输出：6
// 解释： 第 0 天需要展台 a、c、d； 第 1 天需要展台 b、e、d； 第 2 天需要展台 a、c、c、d； 因此，后勤部准备 abccde 的展台，可以满足每天的展览需求;

// 用两个map存储数据，一个全局map存储对应字符的出现的最大次数，一个局部map用来更新每个字符串里出现的字符的次数，
// 如果局部map里的数据大于全局map里的数据，就更新（如果map的key没出现，那他就为0）顺便对sult进行+1，以此来保证全局map是最新、出现次数最大的
func minNumBooths(demand []string) int {
	// 全局map存储出现最多次数的数据
	var dmm = make(map[rune]int, 0)
	// 存储结果
	var sult int
	for _, v := range demand {
		// 局部map用来存储每个字符串里的字符出现的次数
		var tempdm map[rune]int = make(map[rune]int, 0)
		for _, m := range v {
			// 出现就+1
			tempdm[m]++
			// 局部map大于全局map就更新
			if tempdm[m] > dmm[m] {
				sult++
				dmm[m] = tempdm[m]
			}
		}
	}
	return sult
}
