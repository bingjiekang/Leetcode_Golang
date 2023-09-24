package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// st := []int{1, 3, 4, 6, 8}
	// fmt.Println(st, unsafe.Pointer(&st))
	// test2(st)
	// fmt.Println(st, unsafe.Pointer(&st))
}

func sortedSquares(nums []int) (sult []int) {
	length := len(nums)
	for left, right := 0, length-1; left <= right; {
		lf, rg := nums[left]*nums[left], nums[right]*nums[right]
		if lf >= rg {
			sult = append([]int{lf}, sult...)
			// sult = append(sult,)
			left++
		} else {
			sult = append([]int{rg}, sult...)
			right--
		}
	}
	return sult
}
