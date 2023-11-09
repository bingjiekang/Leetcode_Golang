package face10

// 在一个整数数组中，“峰”是大于或等于相邻整数的元素，相应地，“谷”是小于或等于相邻整数的元素。例如，在数组{5, 8, 4, 2, 3, 4, 6}中，{8, 6}是峰， {5, 2}是谷。现在给定一个整数数组，将该数组按峰与谷的交替顺序排序。

// 示例:

// 输入: [5, 3, 1, 2, 3]
// 输出: [5, 1, 3, 2, 3]

// 解析:排序后,降序排列,从下标为1开始对调相邻数既可,因为已经降序,所以后面的数调换到前面后,这个数一定为谷,依次往后效果一样.总体方法:降序排序,调换相邻数

func wiggleSort(nums []int) {
	// 获得nums的长度
	length := len(nums)
	// 降序排列
	quicksort(0, length-1, nums)
	// 从下标1开始,对调相邻数既可
	for temp := 2; temp < length; {
		nums[temp-1], nums[temp] = nums[temp], nums[temp-1]
		temp += 2
	}
}

// 实现快排
func quicksort(begin, end int, nums []int) {
	left, right := begin, end
	if left > right {
		return
	}
	temp := left
	for left < right {
		for left < right && nums[right] <= nums[temp] {
			right--
		}
		for left < right && nums[left] >= nums[temp] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	if left == right {
		nums[temp], nums[left] = nums[left], nums[temp]
	}
	quicksort(begin, left-1, nums)
	quicksort(right+1, end, nums)
}
