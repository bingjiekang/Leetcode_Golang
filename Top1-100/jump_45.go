package main

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

// 示例 1:
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

// 解析:从后往前,因为是求到最后一位最少的步数,因此遍历时从前往后,找到一个能到达或超过最后一位的数的下标,此时,这个数肯定是步骤的最后一跳,依次,将这一位定做最后一位,继续从前向后寻找,找到能到达这个数的下标

// 我们的目标是到达数组的最后一个位置，因此我们可以考虑最后一步跳跃前所在的位置，该位置通过跳跃能够到达最后一个位置。如果有多个位置通过跳跃都能够到达最后一个位置，我们可以「贪心」地选择距离最后一个位置最远的那个位置，也就是对应下标最小的那个位置。因此，我们可以从左到右遍历数组，选择第一个满足要求的位置。找到最后一步跳跃前所在的位置之后，我们继续贪心地寻找倒数第二步跳跃前所在的位置，以此类推，直到找到数组的开始位置。

func jump1(nums []int) int {
	// 贪心算法，从后向前
	length := len(nums) - 1
	// 用来存储次数
	sult := 0
	// 当找到的位置不为第一位时,继续遍历
	for length > 0 {
		// 用来查找能到达或者超过"最后一位"的下标
		for i := 0; i < length; i++ {
			// 得到当前值所达到的最大下标
			target := i + nums[i]
			// 如果这个下标超过"最后一位",则将这一位赋值成最后一位
			if target >= length {
				length = i
				sult++
				break
			}
		}
	}
	// 返回既可
	return sult

}

// 从前往后遍历,每一次遍历获得当前值所能达到的最远下标,如果这个最远下标比前一次遍历的最远下标大,则更新存储这个下标.每一次经过最远边界将sult次数加1,对最远下标进行赋值,小细节:第一跳从下标0开始,最开始边界指向0,这样就不用考虑最后边界的下标大于或者等于最后一位了

//如果我们「贪心」地进行正向查找，每次找到可到达的最远位置，就可以在线性时间内得到最少的跳跃次数。
// 例如，对于数组 [2,3,1,2,4,2,3]，初始位置是下标 0，从下标 0 出发，最远可到达下标 2。下标 0 可到达的位置中，下标 1 的值是 3，从下标 1 出发可以达到更远的位置，因此第一步到达下标 1。
// 从下标 1 出发，最远可到达下标 4。下标 1 可到达的位置中，下标 4 的值是 4 ，从下标 4 出发可以达到更远的位置，因此第二步到达下标 4。
// 在具体的实现中，我们维护当前能够到达的最大下标位置，记为边界。我们从左到右遍历数组，到达边界时，更新边界并将跳跃次数增加 1。
// 在遍历数组时，我们不访问最后一个元素，这是因为在访问最后一个元素之前，我们的边界一定大于等于最后一个位置，否则就无法跳到最后一个位置了。如果访问最后一个元素，在边界正好为最后一个位置的情况下，我们会增加一次「不必要的跳跃次数」，因此我们不必访问最后一个元素。

func jump2(nums []int) int {
	// 从前往后
	// 获取数组长度
	length := len(nums)
	// 边界end
	end := 0
	// maxposition,当前所能到达的最大边界
	maxposition := end
	// 跳跃次数
	sult := 0
	// 遍历下标直到最后一位的前一位
	for i := 0; i < length-1; i++ {
		// 如果当前所能到达的下标大于之前所能到达的最大边界
		if i+nums[i] > maxposition {
			maxposition = i + nums[i] // 赋值最大边界为当前所能达到的最大下标
		}
		if i == end { // 如果当前遍历到边界end,则对次数加1,对边界进行扩大,到当前所能达到的最大边界
			sult++
			end = maxposition
		}
	}
	// 返回结果即可
	return sult
}