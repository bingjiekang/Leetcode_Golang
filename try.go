package main

import (
	"fmt"
)

func main() {
	var tm []int = []int{1, 2, 3, 1}
	fmt.Println(containsNearbyDuplicate(tm, 3))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var tmp map[int]int = make(map[int]int, 1)
	for i, v := range nums {
		if tmp[v] > 0 {
			tmp[v] = i - tmp[v]
			if tmp[v] <= k {
				return true
			}
		} else {
			tmp[v] = i + 1
		}
	}
	return false
}
