package main

// 给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

// 解析:使用hash表来存储遍历其中一个表的数据,将对应的值key每出现一个相同的,就加1,然后遍历另外一个表2,每存在一个相同的,便减1,并把这个加入到返回数组中(优化:存储的时候遍历长度最小的数组,因为是求交集,所以最多相同的个数是长度最短数组的全部遍历)

func intersect(nums1 []int, nums2 []int) (nums []int) {
	// 定义一个hash数组用来存储出现的次数
	var hash map[int]int = make(map[int]int)
	// 获取两个数组的长度
	lnums1, lnums2 := len(nums1), len(nums2)
	// 谁短遍历存储谁
	if lnums1 <= lnums2 {
		// 遍历存储
		for _, v := range nums1 {
			hash[v] += 1
		}
		// 对存储的hash表进行遍历读取,减1并加入到待返回数组
		for _, v := range nums2 {
			// 用来判断v是否存在
			sult, err := hash[v]
			// 如果存在且对应的值还不为0 就可以加入并且减1
			if err && sult > 0 {
				hash[v] -= 1
				nums = append(nums, v)
			}
		}
	} else { // 和上方方法一样,对调一下数组即可
		for _, v := range nums2 {
			hash[v] += 1
		}
		for _, v := range nums1 {
			sult, err := hash[v]
			if err && sult > 0 {
				hash[v] -= 1
				nums = append(nums, v)
			}
		}
	}
	// 返回即可
	return
}
