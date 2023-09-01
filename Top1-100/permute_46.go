package main

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// 解析:将数组分为两部分,一部分为以及加入到数组中的,另一部分为待加入的,通过回溯法进行解决,通过对index的回溯,如果index指向最后一位则加入这个数组,否则就让i从index遍历到range,加入到回溯算法中,实现每一层的 1-1/2/3/4 2-2/3/4 3-3/4 4-4交换,

func permute(nums []int) (sult [][]int) {
	// 获取长度
	length := len(nums)
	// rsv函数实现回溯对调
	var rsv func(index int)
	rsv = func(index int) {
		// 如果所有数都填完,则加入到数组里,切片是浅拷贝,使用copy函数进行深拷贝,需要指定对应的len
		if index == length-1 {
			var temp []int = make([]int, length)
			copy(temp, nums)
			sult = append(sult, temp)
		}
		// 实现回溯和对调
		for i := index; i < length; i++ {
			// 动态维护并对调数据
			nums[i], nums[index] = nums[index], nums[i]
			// 深入每一层,对调对应数据
			rsv(index + 1)
			// 将数据归为未对调数据
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	// 开始即可
	rsv(0)
	return
}
