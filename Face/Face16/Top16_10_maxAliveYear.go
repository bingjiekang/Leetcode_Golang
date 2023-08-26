package main

// 给定 N 个人的出生年份和死亡年份，第 i 个人的出生年份为 birth[i]，死亡年份为 death[i]，实现一个方法以计算生存人数最多的年份。

// 你可以假设所有人都出生于 1900 年至 2000 年（含 1900 和 2000 ）之间。如果一个人在某一年的任意时期处于生存状态，那么他应该被纳入那一年的统计中。例如，生于 1908 年、死于 1909 年的人应当被列入 1908 年和 1909 年的计数。

// 如果有多个年份生存人数相同且均为最大值，输出其中最小的年份。

// 示例：

// 输入：
// birth = [1900, 1901, 1950]
// death = [1948, 1951, 2000]
// 输出： 1901

// 解析:题目表示所有人出生在1900到2000年,则可以建立一个长度为101的切片存储对应年份,将对应下标进行hash存储,即:对1900取模,将对应值下标对应的值进行加1操作,从birth和death,对每个对应的下标进行加1操作,同时记录当前年份最大人次,和对应最佳年份,最后遍历完后返回最大人次对应的最佳年份即可

func maxAliveYear(birth []int, death []int) int {
	// base用来存储固定值用来进行hash取值
	const base int = 1900
	// 定义一个长度为101的切片长度
	var lt []int = make([]int, 101)
	// maxvalue存储人次,mdata存储当前年份
	var maxvalue, mdata int
	// 获取人数的
	length := len(birth)
	// 从第一个开始遍历,遍历到结尾,对birth和death切片同时遍历
	for i := 0; i < length; i++ {
		// 将对应的年份进行hash取值,得到对应下标,start用来从birth遍历到death结尾
		for start, end := birth[i]%base, death[i]%base; start <= end && start < 101; start++ {
			// 对 相应值进行加1操作
			lt[start]++
			// 如果当前人次大于最大人次 则进行更新
			if lt[start] > maxvalue {
				maxvalue = lt[start]
				mdata = 1900 + start
			} else if lt[start] == maxvalue && mdata > 1900+start { // 如果当前人次等于最大人次,则比较当前年份是否小于最佳年份,如果最佳年份大于当前年份,则对年份进行更新
				mdata = 1900 + start
			}
		}
	}
	// 返回最佳年份即可
	return mdata

}
