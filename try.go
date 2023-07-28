package main

import (
	"fmt"
	"sort"
)

// "strconv"
// "time"
// "math"

func main() {

	// n := 3

	// for i := 0; i <= 4; i++ {
	// 	switch i {
	// 	case 3:
	// 		fmt.Println(3)
	// 	default:
	// 		fmt.Println(99999)
	// 	}

	// }
	var temp ST
	var nums []int = []int{1, 3, 8, 1, 1, 2, 1, 3, 9, 5, 7, 4}
	temp.nums = nums
	sort.Sort(sm(temp))
	// fmt.Println(maximumGap(nums))
	fmt.Println(temp.nums)

}

type ST struct {
	nums []int
}

type sm ST

func (this sm) Len() int {
	return len(this.nums)
}

func (this sm) Swap(n, m int) {
	this.nums[n], this.nums[m] = this.nums[m], this.nums[n]
}

func (this sm) Less(i, j int) bool {
	return this.nums[i] > this.nums[j]
}
