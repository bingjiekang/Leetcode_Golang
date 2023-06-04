package main

// 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。
// 示例 1:
// 输入: num = 38
// 输出: 2
// 解释: 各位相加的过程为：
// 38 --> 3 + 8 --> 11
// 11 --> 1 + 1 --> 2
// 由于 2 是一位数，所以返回 2。
// 各个位数相加直到最后只有一位数,这位数是数根,除9个倍数外,任何一个数对9取模后的余数都为数根,为了让9的倍数也适用,则(num-1)%9+1即可,对原数-1后,取模后再-1

func addDigits(num int) int {
	temp := dp(num)
	for temp >= 10 {
		temp = dp(temp)
	}
	return temp
	// 第二种方法,根据数学方法求数根
	//  return (num-1)%9 + 1
}

func dp(num int) int {
	sult := 0
	for num > 0 {
		sult = num%10 + sult
		num /= 10
	}
	return sult
}
