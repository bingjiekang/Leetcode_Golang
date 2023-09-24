package main

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
// 示例 1：

// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]
// 排序后，数组变为 [0,1,9,16,100]

// 双指针:左指针从前往后,右指针从后往前,谁大谁加入sult结果切片中,最后遍历完后,新切片sult是从大到小降序的,再倒序转换一下就可以了

func sortedSquares(nums []int) (sult []int) {
	// 获取长度
	length := len(nums)
	// left从前往后,右指针从后往前,
	for left, right := 0, length-1; left <= right; {
		// 获得对应数据的平方并比较
		lf, rg := nums[left]*nums[left], nums[right]*nums[right]
		// 谁大谁放入sult中
		if lf >= rg {
			sult = append(sult, lf)
			left++
		} else {
			sult = append(sult, rg)
			right--
		}
	}
	// 将切片翻转一下即可
	for i, j := 0, length-1; i < j; {
		sult[i], sult[j] = sult[j], sult[i]
		i++
		j--
	}
	return sult
}
