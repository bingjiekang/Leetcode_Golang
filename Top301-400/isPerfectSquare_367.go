package main

// 给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
// 完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。
// 不能使用任何内置的库函数，如  sqrt

// 解析:利用二分法进行求取完全平方数,反复二分,直到mid*mid等于结果,如果最后left>right则返回false,如果找到则返回true

func isPerfectSquare(num int) bool {
	// 规划左边界和有边界
	left, right := 1, num
	// 当left小于等于右边时,一直进行二分
	for left <= right {
		mid := (left + right) / 2
		// 当mid*mid小于num时,将mid赋值left
		if mid*mid < num {
			left = mid + 1
		} else if mid*mid > num { // 当mid*mid大于num时,将mid赋值right
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
