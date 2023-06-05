package main

// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
// 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

// 解析:采用二分法找到这个数即可

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	// 左右left和right
	left, right := 1, n
	var mid int
	for left <= right {
		mid = (left + right) / 2
		// 用来接收此版本是否出错
		mid0 := isBadVersion(mid - 1)
		mid1 := isBadVersion(mid)
		mid2 := isBadVersion(mid + 1)
		// 当前版本出错,且前一版本也出错,缩小范围至左部
		if mid1 == true && mid0 != false {
			right = mid
		} else if mid1 == false && mid2 != true { // 当前版本不出错,且下一版不出错,缩小范围至右部
			left = mid
		} else { // 否则就是当前版本内已经有第一个出错的节点
			if mid1 == false { // 证明是下一个节点为第一个出错的点
				mid = mid + 1
			} // 否则就是本节点
			break
		}
	}
	return mid
}
