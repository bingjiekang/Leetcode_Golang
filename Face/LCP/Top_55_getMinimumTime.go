package main

// 欢迎各位勇者来到力扣新手村，本次训练内容为「采集果实」。
// 在新手村中，各位勇者需要采集一些果实来制作药剂。time[i] 表示勇者每次采集 1～limit 颗第 i 种类型的果实需要的时间（即每次最多可以采集 limit 颗果实）。
// 当前勇者需要完成「采集若干批果实」的任务， fruits[j] = [type, num] 表示第 j 批需要采集 num 颗 type 类型的果实。采集规则如下：
// 按 fruits 给定的顺序依次采集每一批次
// 采集完当前批次的果实才能开始采集下一批次
// 勇者完成当前批次的采集后将清空背包（即多余的果实将清空）
// 请计算并返回勇者完成采集任务最少需要的时间。

// 示例 1：
// 输入：time = [2,3,2], fruits = [[0,2],[1,4],[2,1]], limit = 3
// 输出：10

// 解释： 由于单次最多采集 3 颗 第 0 批需要采集 2 颗第 0 类型果实，需要采集 1 次，耗时为 2*1=2 第 1 批需要采集 4 颗第 1 类型果实，需要采集 2 次，耗时为 3*2=6 第 2 批需要采集 1 颗第 2 类型果实，需要采集 1 次，耗时为 2*1=2 返回总耗时 2+6+2=10

// 遍历数组fruits,判断是否可以以整数形式采摘完对应水果,若果不能,则对采摘次数加1,将对应的采摘时间和次数结合起来相加即可

func getMinimumTime(time []int, fruits [][]int, limit int) int {
	// cout为对应水果采摘次数,tm为采摘全部水果需要的时间
	var cout, tm int
	// 遍历数组fruits
	for _, v := range fruits {
		// 如果可以以整数次采摘完毕,则计算得出对应次数
		if v[1]%limit == 0 {
			cout = v[1] / limit
		} else { // 如果不可以,则对应次数加1
			cout = v[1]/limit + 1
		}
		// 计算需要时间
		tm += cout * time[v[0]]
		// 初始化cout
		cout = 0
	}
	// 返回即可
	return tm
}