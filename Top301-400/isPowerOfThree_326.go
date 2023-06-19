package main

// 给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
// 整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3^x

// 解析:反复除3 直到余数不为0 得到的结果若不为1则返回false,否则返回true

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	} else {
		for n%3 == 0 {
			n /= 3
		}
		return n == 1
	}
}
