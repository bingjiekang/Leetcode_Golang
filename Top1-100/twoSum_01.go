package main

/*
要求:给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。你可以按任意顺序返回答案。
*/

func twoSum(nums []int, target int) (list []int) {
	// 利用hash表存储 将列表里的值对应key，序号对应value
	var HashMap map[int]int = make(map[int]int, 0)
	for i, v := range nums {
		// map返回两个参数 第一个是value，第二个是 是否存在所找key
		if _, ok := HashMap[target-v]; ok { // 如果target-当前值 在hash表里，证明之前已经出现
			list = append(list, HashMap[target-v], i)
			return list
		} else { // 如果当前值不在hash表里，将该值加入hash表
			HashMap[v] = i
		}
	}
	return nil
}
