package top1100

import "sort"

// 56. 合并区间
// Type : 中等
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

// 示例 1：
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

// 对二维数组里面的左端点进行排序，如果两个数组能合并，则其左端点一定相邻，如果一个数数组的左端点大于一个数组里的右端点，则其不能合并
// 如果当前区间的左端点在数组 merged 中最后一个区间的右端点之后，那么它们不会重合，我们可以直接将这个区间加入数组 merged 的末尾；
// 否则，它们重合，我们需要用当前区间的右端点更新数组 merged 中最后一个区间的右端点，将其置为二者的较
func merge(intervals [][]int) (sult [][]int) {
	// 定义待排序的用来存储数据的结构体切片
	var newNums []Nums
	// 临时变量，用来存储左右断电及赋值
	var temp Nums
	for _, v := range intervals {
		temp.Start = v[0]
		temp.End = v[1]
		newNums = append(newNums, temp)
	}
	// 对结构体切片进行排序
	sort.Sort(md(newNums))
	// 用来记录sult数组的最后一个数组元素的下标
	cout := -1
	for _, v := range newNums {
		// sult 数组里没有则直接加入
		if cout < 0 {
			sult = append(sult, []int{v.Start, v.End})
			cout++
		} else {
			//  如果当前数组的左端点大于sult数组里最后数组元素的右端点则直接加入
			if v.Start > sult[cout][1] {
				sult = append(sult, []int{v.Start, v.End})
				cout++
			} else { // 否则则更新原最后数组元素的end 取最大值即可
				sult[cout][1] = max(sult[cout][1], v.End)
			}
		}
	}
	return
}

// 取最大值
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 以下用来对数组进行指定规则排序
type Nums struct {
	Start int
	End   int
}

// md 实现sort.interface 的Len、Swap、Less方法，即可调用sort.Sort()记得里面需要用md包括
type md []Nums

func (m md) Len() int {
	return len(m)
}

// 重点：根据谁进行比较，这里是根据数组的左端点进行比较
func (m md) Less(i, j int) bool {
	return m[i].Start < m[j].Start
}

func (m md) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
