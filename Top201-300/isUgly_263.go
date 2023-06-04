package main

// 丑数 就是只包含质因数 2、3 和 5 的正整数。
// 给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。

// 解析:用该数反复除这三个数,如果最后结果为1,则证明这个数除了这三个数外,没有其他质因数.n = 2^a+3^b+5^c
func isUgly(n int) bool {
	// 用于遍历的数组
	var lt []int = []int{2, 3, 5}
	// n小于0不会只包含质因数,因为他们都是正整数
	if n <= 0 {
		return false
	}
	// 外部循环遍历2,3,5
	for _, v := range lt {
		// 内部循环对n进行分解
		for n%v == 0 {
			n /= v
		}
	}
	// 如果最后结果为1则返回true否则返回false
	return n == 1
}
