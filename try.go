package main

import (
	"fmt"
	// "strconv"
)

// "strconv"
// "time"
// "math"

func main() {

	// var lt []int = []int{1, 2, 3, 4, 2, 2, 1, 2, 8, 6, 6, 6, 4, 8, 9}
	// QuickSort(lt, 0, len(lt)-1)
	// fmt.Print(lt)
	// x := fmt.Sprintf("%.0b", 7)
	// fmt.Println(x)
	// x := 1<<3
	// for i,v := range {

	// }
	// t := byte(2)
	// fmt.Println(t)
}

var nm int = len("2344")

func jump(nums []int) int {
	dp := make([][]int, 10000, 10000)
	// var dp [10000][10000]int = make([10000][10000]int, 0)
	length := len(nums)
	for i := length - 1; i >= 0; i-- {
		dp[i] = make([]int, 10000)
		for j := i; j < length; j++ {
			if i == j {
				dp[i][j] = 0
			} else if nums[i] >= j-i {
				dp[i][j] = 1
			} else {
				min := 10000
				for temp := i + 1; temp < j; temp++ {
					min_temp := dp[i][temp] + dp[temp][j]
					if min_temp < min {
						min = min_temp
					}
				}
				dp[i][j] = min
			}
		}
	}
	return dp[0][length-1]
}
