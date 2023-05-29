package main

// 要求:给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

// 解析:维护一个哈希表,由于是如果两个相同的数之间距离小于等于k值,返回true,全部大于k的或者没有重复的返回false,则建立哈希表,如果该数据没有出现,则把他加入到表中,如果出现,则判断当前数据和表中数据的value的距离是否小于k,如果小于k则返回value,否则更新这个数据对应的value为当前值的下标,如果全部遍历完都不满足,返回false

func containsNearbyDuplicate(nums []int, k int) bool {
	// 建立一个hash表
	var tmp map[int]int = make(map[int]int, 1)
	// 遍历
	for i, v := range nums {
		// 用来判断数据v是否已经在hash表中,ok为true表示在,为false表示不在
		_, ok := tmp[v]
		if ok {
			// 该值已经存在,得到他们之间的距离
			tmp[v] = i - tmp[v]
			// 如果小于k则返回true
			if tmp[v] <= k {
				return true
			} else { //否则当前值为新的下标,表示在此之前没有满足条件的
				tmp[v] = i
			}
		} else { // 如果hash表中不存在,则加入新的值
			tmp[v] = i
		}
	}
	return false
}
