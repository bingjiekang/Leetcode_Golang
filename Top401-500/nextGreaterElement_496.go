package top401500

// 496. 下一个更大元素 I
// Type: 简单
// nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
// 给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
// 对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
// 返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。

// 示例 1：
// 输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
// 输出：[-1,3,-1]
// 解释：nums1 中每个值的下一个更大元素如下所述：
// - 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
// - 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
// - 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。

// 遍历nums2记录对应数的下标，由于nums1是nums2的子集，遍历nums1找到对应nums2的位置遍历后面是否有满足条件的，如果满足则加入对应数据，没有就加入-1，使用target标记是否进入循环
func nextGreaterElement(nums1 []int, nums2 []int) (sult []int) {
	length := len(nums2)
	// 记录nums2对应数和其下标
	var dt map[int]int = make(map[int]int, 0)
	// 用来标记是否有数满足条件
	var target bool
	// 遍历nums2并记录对应数和其下标
	for k, v := range nums2 {
		dt[v] = k
	}
	// 遍历nums1找到满足条件的对应数
	for _, v := range nums1 {
		target = false
		for i := dt[v]; i < length; i++ {
			if nums2[i] > v {
				target = true
				sult = append(sult, nums2[i])
				break
			}
		}
		if !target {
			sult = append(sult, -1)
		}
	}
	return
}
