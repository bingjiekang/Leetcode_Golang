package main

import (
	"fmt"
)

// "strconv"
// "time"
// "math"

func main() {

	fmt.Println((1 + 1) / 2)

}

func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left < right {
		mid := (left + right) / 2
		if mid*mid < num {
			left = mid + 1
		} else if mid*mid > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
