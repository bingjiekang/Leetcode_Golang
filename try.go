func wiggleSort(nums []int) {
	length := len(nums)
	quicksort(0, length-1, nums)
	for temp := 2; temp < length; {
		nums[temp-1], nums[temp] = nums[temp], nums[temp-1]
		temp += 2
	}
}

func quicksort(begin, end int, nums []int) {
	left, right := begin, end
	if left > right {
		return
	}
	temp := left
	for left < right {
		for left < right && nums[right] <= nums[temp] {
			right--
		}
		for left < right && nums[left] >= nums[temp] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	if left == right {
		nums[temp], nums[left] = nums[left], nums[temp]
	}
	quicksort(begin, left-1, nums)
	quicksort(right+1, end, nums)
}