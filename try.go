package main

import (
	"fmt"
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
	// num1 := "3"
	// s1 := []rune(num1)
	// num2 := "5"
	// s2 := []rune(num2)
	// fmt.Println(s1 * s2)

}

func singleNumber(nums []int) int {
	var temp map[int]int = make(map[int]int, 0)
	for _, v := range nums {
		temp[v]++
		delete(temp, v)
	}
	for i, v := range temp {
		if v == 1 {
			return i
		}
	}
	return 0
}
