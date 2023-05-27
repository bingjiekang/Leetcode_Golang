package main

// 例题:编写一个算法来判断一个数 n 是不是快乐数。
// 快乐数」 定义为：
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false

// 解析:通过此方法,最终有三种情况,1:出现循环,则证明一定不是快乐数,2:最后化为1,是快乐数,3:无限大(不成立),因为每个数的最大数字为9,对应的每位9/99/999/9999,经过此方法后,都会化为比当前数小的数字,则证明此方法算出的数不会一直增大

func isHappy(n int) bool {
	// 定义哈希表用来存储每一个需要判断的数据
	var hash map[int]bool = make(map[int]bool, 1)
	// 如果最终不能重复到1,那么一定会有循环,即一定在hash里存储过,返回false,
	for n != 1 {
		if hash[n] {
			return false
		} else {
			hash[n] = true
		}
		n = AllAdd(n)
	}
	// 最终结果重复到1,则是快乐数
	return true
}

// 用来分解每位数并获得结果例如:19-> 1^2 + 9^2 = 82
func AllAdd(m int) int {
	var temp, result int
	for m > 0 {
		temp = m % 10
		result += temp * temp
		m /= 10
	}
	return result
}
