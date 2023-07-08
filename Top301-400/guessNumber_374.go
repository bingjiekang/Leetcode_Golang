package main

// 猜数字游戏的规则如下：
// 每轮游戏，我都会从 1 到 n 随机选择一个数字。 请你猜选出的是哪个数字。
// 如果你猜错了，我会告诉你，你猜测的数字比我选出的数字是大了还是小了。
// 你可以通过调用一个预先定义好的接口 int guess(int num) 来获取猜测结果，返回值一共有 3 种可能的情况（-1，1 或 0）：
// -1：我选出的数字比你猜的数字小 pick < num
// 1：我选出的数字比你猜的数字大 pick > num
// 0：我选出的数字和你猜的数字一样。恭喜！你猜对了！pick == num
// 返回我选出的数字。

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 *
 */
// 解析:利用二分法进行查找和对比,guess是用来比较提供的函数的
func guess(num int) int

func guessNumber(n int) (sult int) {
	// left和right同时确定范围
	left, right := 1, n
	// 当left小于right时
	for left <= right {
		//二分法确定中间的位置
		mid := (left + right) / 2
		// 如果猜测的数字大于待定数字,指定右边界
		if guess(mid) < 0 {
			right = mid - 1
		} else if guess(mid) > 0 { // 指定左边界
			left = mid + 1
		} else { // 返回结果
			sult = mid
			break
		}
	}
	return sult
}
