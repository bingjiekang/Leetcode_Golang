package main

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 解析:左指针左边均为非零数；右指针左边直到左指针处均为零。因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。
func moveZeroes(nums []int) {
	// //right为不为0的数，将不为0的数和left指向的0数交换即可
	// left, right, n := 0, 0, len(nums)
	// for right < n {
	// 	//如果right 不为0则交换,为0则加1
	//     if nums[right] != 0 {
	// 		nums[left], nums[right] = nums[right], nums[left]
	// 		//保证对应的left即使不为0也只会和自身交换一次
	// 		left++
	//     }
	//     right++
	// }
	// 方法一样
	var left, right int
	length := len(nums)
	for left <= right && right < length {
		for left < length && nums[left] != 0 {
			left++
		}
		for right = left + 1; right < length && nums[right] == 0; {
			right++
		}
		if right < length {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right++
		} else {
			break
		}
	}

}
