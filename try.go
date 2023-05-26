package main

import (
	"fmt"
)

func main() {
	var tm []int = []int{1, 2, 3, 1}
	fmt.Println(massage(tm))
}

func massage(nums []int) int {
	result := 0
	dp0 := 0
	dp1 := nums[0]
	for i, v := range nums {
		if i > 0 {
			result = max(max(dp0, dp1), dp0+v) //2 4 4
			dp0 = max(dp0, dp1)                //1 2 4
			dp1 = dp0 + v                      // 2 4 3
		}
	}
	return result
}

// max函数
func max(n int, m int) int {

	if n > m {
		return n
	}
	return m

}
