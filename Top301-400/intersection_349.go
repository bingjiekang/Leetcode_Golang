package main

// 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。
// 我们可以 不考虑输出结果的顺序 。
// 解析：遍历一个数组后，将所有的存在的加入到hash表里，遍历另一个数组，
// 若同时存在则加入结果列表中，并将此key对应的value设为false，否则就继续遍历

func intersection(nums1 []int, nums2 []int) (nums []int) {
	// 定义一个bool形的hash表
	var temp map[int]bool
	// 声明并make一下
	temp = make(map[int]bool, 1)
	// 遍历第一个数组
	for _, v := range nums1 {
		if !temp[v] {
			temp[v] = true
		}
	}
	// 遍历第二个数组
	for _, v := range nums2 {
		if temp[v] {
			// 如果存在就加入并将hash表的值置为false，否则其他就继续遍历
			nums = append(nums, v)
			temp[v] = false
		}
	}
	return
}
