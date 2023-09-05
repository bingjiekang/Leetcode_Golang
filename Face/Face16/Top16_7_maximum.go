package main

// 编写一个方法，找出两个数字a和b中最大的那一个。不得使用if-else或其他比较运算符。

// 示例：

// 输入： a = 1, b = 2
// 输出： 2

// 使用位运算,移动位,从移动31位开始,到移动到0位,比较每一位谁大,返回谁,一样大的话,最终返回任意数即可
func maximum(a int, b int) int {
	// >>向右移位 从31开始到0
	for i := 31; i >= 0; i-- {
		if a>>i > b>>i {
			return a
		} else if a>>i < b>>i {
			return b
		}
	}
	return a
}
