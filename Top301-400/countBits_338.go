package main

// 给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
// 通过反复除2和对2求余得到1的个数
func countBits(n int) (sult []int) {
	// 对判断的结果进行加1处理
	for i := 0; i <= n; i++ {
		sult = append(sult, CountBit(i))
	}
	return
}

func CountBit(n int) (sum int) {
	// 如果n大于2或者n == 1
	for n/2 != 0 || n == 1 {
		// 对2求模为1则2进制含1,不为1则不含1
		if n%2 == 1 {
			sum++
		}
		// 反复除2
		n /= 2
	}
	return
}
