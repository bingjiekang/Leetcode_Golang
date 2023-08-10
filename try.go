package main

import (
	"fmt"
)

func main() {
	capacity := []int{54, 18, 91, 49, 51, 45, 58, 54, 47, 91, 90, 20, 85, 20, 90, 49, 10, 84, 59, 29, 40, 9, 100, 1, 64, 71, 30, 46, 91}
	rocks := []int{14, 13, 16, 44, 8, 20, 51, 15, 46, 76, 51, 20, 77, 13, 14, 35, 6, 34, 34, 13, 3, 8, 1, 1, 61, 5, 2, 15, 18}
	length := len(capacity)
	quicksort(0, length-1, capacity)
	quicksort(0, length-1, rocks)
	fmt.Println(capacity, rocks)
}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	// 获得每个背包空的位置，然后排序，从小到大计数直到数量大于额外石头数量
	length := len(capacity)
	// ord := []int{}
	var cout, sult int
	for i := 0; i < length; i++ {
		rocks[i] = capacity[i] - rocks[i]
	}
	quicksort(0, length-1, rocks)
	for _, v := range rocks {
		sult += v
		if sult <= additionalRocks {
			cout++
		} else {
			break
		}
	}
	return cout

}

func quicksort(start, end int, nums []int) {
	if start > end {
		return
	}
	left, right := start, end
	temp := start
	for left < right {
		for left < right && nums[right] >= nums[temp] {
			right--
		}
		for left < right && nums[left] <= nums[temp] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	if left == right {
		nums[temp], nums[left] = nums[left], nums[temp]
	}
	quicksort(start, left-1, nums)
	quicksort(right+1, end, nums)
}
