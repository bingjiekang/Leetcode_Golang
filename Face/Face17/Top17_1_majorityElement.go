package face17

// 数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。

// 示例 1：

// 输入：[1,2,5,9,5,9,5,5,5]
// 输出：5

// 示例 2：

// 输入：[3,2]
// 输出：-1

func majorityElement(nums []int) int {
	// // 定义一个map存储元素出现的次数
	// var hash map[int]int = make(map[int]int,0)
	// // length 存储一半的长度
	// length := len(nums)/2
	// // 遍历数组，记录数字出现的次数
	// for _,v := range nums{
	//     hash[v]++
	// }
	// // 遍历map，如果有大于一半的数组值则返回，否则返回-1
	// for k,v := range hash{
	//     if v>length{
	//         return k
	//     }
	// }
	// return -1

	// 投票法
	// 由于求出现超过一半长度的数，遍历每个数，记录这个数num和出现的次数，如果count为0，则赋值这个数为num，如果num等于这个数，则count++，不等于则cout--，由于如果存在超过一半长度的数，那么这个数一定会和其他数消完后还余1，如果为空则不存在，还有一种情况是奇数次且不存在超过数组长度一半的数，所以需要第二次遍历查看这个数是否合理
	// num记录这个数，cout记录出现的次数，length记录数组长度的一半
	var num, cout, length int
	length = len(nums) / 2
	// 遍历这个数组，出现次数为0，赋值num和对cout++,大于0则判断是否相同相同cout++，不同cout--
	for _, v := range nums {
		if cout == 0 {
			num = v
			cout++
		} else {
			if num == v {
				cout++
			} else {
				cout--
			}
		}
	}
	// 出现次数大于0，判断这个数出现的次数是否真的超过数组的一半
	if cout > 0 {
		cout = 0
		for _, v := range nums {
			if v == num {
				cout++
			}
		}
		if cout > length {
			return num
		}
	}
	return -1

}
