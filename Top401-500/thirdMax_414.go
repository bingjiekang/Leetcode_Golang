package main

// 给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。

// 解析:使用快排对数组进行排序,然后去重得到不重复的数组,得到倒数第三位既可,如果不足三位返回最大的

func thirdMax(nums []int) int {
	// 先快排
	QuickSort(nums, 0, len(nums)-1)
	// 再去重
	nums = Wreduce(nums)
	// 得到去重后的数组长度
	length := len(nums)
	// 长度小于3 返回最大值
	if length < 3 {
		return nums[length-1]
	} else { // 否则返回倒数第三位
		return nums[length-3]
	}
}

// [快速排序]传入切片 原地修改切片数据
func QuickSort(nums []int, left int, right int) {
	// 存储原来left和right的位置
	l, r := left, right
	// 如果left和right指针指向同一位置或者交叉返回
	if left >= right {
		return
	}
	// 假设待比较的值,在每个数组的第一个
	temp := left
	// 左指针小于右指针的时候,继续比较和交换
	for left < right {
		// 先从右往左找到第一个小于 nums[temp]假定值 的值
		for left < right && nums[right] >= nums[temp] {
			right--
		}
		// 再从左往右找到第一个大于 nums[temp]假定值 的值
		for left < right && nums[left] <= nums[temp] {
			left++
		}
		// 交换左右这两个值
		nums[left], nums[right] = nums[right], nums[left]
	}
	// 此时left和right指向同一个数,将他们指向的数和最开始假定的值交换即可
	nums[temp], nums[left] = nums[left], nums[temp]
	// 然后递归遍历排序左边的数据
	QuickSort(nums, l, left-1)
	// 递归遍历排序右边的数据
	QuickSort(nums, right+1, r)
}

// 去重
func Wreduce(nums []int) []int {
	// 获取数组的长度
	length := len(nums)
	// 指针指向前两位数
	left, right := 0, 1
	// 当右指针小于长度时
	for right < length {
		// 如果左指针不等于右指针,左指针右移,指针指向的数赋值为右指针指向的数
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
		// 右指针每遍历一次向右移动一次
		right++
	}
	// 返回去重后的数组
	return nums[:left+1]
}
