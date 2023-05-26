package main

/*
要求:一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
*/
// 方法:采用动态规划:定义 dp[i][0]表示前i个预约,当前i不预约,dp[i][1]表示前i个预约,当前i预约.从前往后,假设已经知道了第dp[i-1]的状态,最优值,则考虑当前的,如果当前这个不预约,则前一个可以预约也可以不预约即:max(dp[i-1][0],dp[i-1][1]),如果当前的预约,则前一个一定不能预约,则当前这个为dp[i-1][0]+当前值

func Massage(nums []int) int {
	// 空列表返回
	if len(nums) == 0 {
		return 0
	}
	// 结果接收
	result := nums[0]
	// 初始化第i-1个不预约,dpi的值
	dp0 := 0
	// 初始化第i-1预约,dpi的值
	dp1 := nums[0]
	for i, v := range nums {
		if i > 0 {
			// 获得当前 不预约的值
			x := Max(dp0, dp1)
			// 当前 预约后的值
			y := dp0 + v
			result = Max(x, y)
			// 更新这些值即可
			dp0 = x
			dp1 = y
		}
	}
	return result
}

// max函数
func Max(n int, m int) int {

	if n >= m {
		return n
	}
	return m

}
