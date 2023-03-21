package main

/*
要求:
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

func majorityElement(nums []int) int {
	// 使用哈希表存储对应数据，对 对应的数字出现的次数进行加一
	// var Hash_map map[int]int = make(map[int]int,0)
	// length := (len(nums))/2
	// for _,v :=range nums{
	//     Hash_map[v]++
	//     if Hash_map[v]>length{
	//         return v
	//     }
	// }
	// return 0
	// 使用Boyer-Moore 投票算法，因为多数元素肯定大于这个数组的一半元素，那么出现相同的加1，不同的减1，因为多数元素多，所以结果肯定为正值
	num, count := 0, 0
	for _, v := range nums {
		// 如果计算为0，将当前值赋成多数元素，并将值（本身）+1
		if count == 0 { // count消完，赋值新数，只要这个数是 多数 结果肯定大于0
			num = v
			count++
		} else if num == v { // 如果出现相同的就+1
			count++
		} else {
			count-- // 不同的就-1
		}
	}
	return num
}
