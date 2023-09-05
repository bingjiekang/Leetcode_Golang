package main

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

// 解析:使用二分法确定大小

func mySqrt(x int) int {
	var sult int
	// left为0,right为对应值,二分确定mid值的平方等于结果,
	for left, right := 0, x; left <= right; {
		mid := (left + right) / 2
		// 如果mid*mid的值小于x则left=mid+1向右查询,并赋值,mid*mid的值等于x则sult最终为mid*mid
		if mid*mid <= x {
			sult = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return sult
}
