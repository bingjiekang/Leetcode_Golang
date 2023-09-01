package main

import (
	"fmt"
)

func main() {
	// var a []int = make([]int, 3)
	// v := []int{1, 2, 3}
	// fmt.Println(a, v)
	// copy(a, v)
	// fmt.Println(a, v)
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) (sult [][]int) {
	// 获取长度
	length := len(nums)
	// 遍历，将切片分为已加入和未加入两组，递归将未加入的加入到即可
	var rsv func(index int)

	rsv = func(index int) {
		if index == length-1 {
			var temp []int = make([]int, length)

			copy(temp, nums)
			sult = append(sult, temp)
			// return
		}

		for i := index; i < length; i++ {
			nums[i], nums[index] = nums[index], nums[i]
			rsv(index + 1)
			nums[i], nums[index] = nums[index], nums[i]
		}
		// return

	}
	rsv(0)

	return

}
