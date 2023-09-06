package main

// 给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。注意：n-m尽量最小，也就是说，找出符合条件的最短序列。函数返回值为[m,n]，若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。

// 示例：

// 输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
// 输出： [3,9]

// 双指针
// 从前往后找到最后一个非递增序列，且其肯定小于右边的递增序列，找到最后一个小于起右边递增序列的第一个数
// 从后往前找到第一个递减序列（最左边从右往左），其肯定大于左边的递增序列（从左往右），找到第一个大于其中左边递增序列的最后一个数
// 例如：1,2,4,7,10,11,7,12,6,7,16,18,19 end从左往右最后找到了 7 （16前面那个数）即下标为9，begin从右往左最后找到了第一个7 即4后面那个， 下标为3 ，
func subSort(array []int) []int {
	length := len(array)
	var end, begin int = -1, -1
	if length == 0 {
		return []int{begin, end}
	}
	var max, min int = array[0], array[length-1]
	for i := 0; i < length; i++ {
		// 从左往右找到最后一个非递增序列的数
		if array[i] > max {
			max = array[i]
		} else if array[i] < max { // 最短的待调整子序列，如果不是最短的可以加上等号
			end = i
		}
		// 从右往左找到第一个非递减序列的数
		bg_i := length - 1 - i
		if array[bg_i] < min {
			min = array[bg_i]
		} else if array[bg_i] > min {
			begin = bg_i
		}
	}
	return []int{begin, end}
}
