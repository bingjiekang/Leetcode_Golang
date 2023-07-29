package main

// 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，x^n）。

// 解析:通过递归,x x^2 x^4 x^8 x^16/x x^2 x^4 x^9 x^19等.通过从右往左,递归得到想要的结果,时间logn,空间logn,判断如果为2的倍数则直接相乘,如果不为2的倍数,则相乘完再乘个自身.

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return dfs(x, n)
	} else {
		return 1 / dfs(x, (-1)*n)
	}

}

// 递归用来得到结果
func dfs(x float64, n int) float64 {
	// 递归边界,任何数的0次幂结果都为1
	if n == 0 {
		return 1
	}
	// 进行递归调用
	sult := dfs(x, n/2)
	// 如果次数是2的倍数则返回平方
	if n%2 == 0 {
		return sult * sult
	} else { // 不是2的倍数返回平方再乘本身
		return sult * sult * x
	}
}
