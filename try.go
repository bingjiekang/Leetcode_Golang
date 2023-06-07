package main

import (
	"fmt"
)

func main() {
	// var tm, tmn *ListNode
	// tm.Val = 1
	// tmn.Val = 1
	// tm.Next = tmn
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(Constructor(nums))
}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	sums = nums
	return NumArray{sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	sult := 0
	for left <= right {
		sult += this.nums[left]
		left++
	}
	return sult
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
