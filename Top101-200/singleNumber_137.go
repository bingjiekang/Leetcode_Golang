package main

// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

// 示例 1：

// 输入：nums = [2,2,3,2]
// 输出：3
// 解析使用map当出现的value值为3时删除这个值,最后遍历完后一定只剩一个 这个一个就是仅有的值

func singleNumber(nums []int) (sult int) {
	// 定义一个map用来存储出现的次数
	var temp map[int]int = make(map[int]int, 0)
	// 遍历数组,如果value值为3则删除这个
	for _, v := range nums {
		temp[v]++
		if temp[v] == 3 {
			delete(temp, v)
		}
	}
	// 仅剩的一个即为结果
	for i, _ := range temp {
		sult = i
	}
	return
}
