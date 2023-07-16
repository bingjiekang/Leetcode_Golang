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

// 解析使用位运算,由于1-n的二进制位对应的就是每个子集出现的位置,则可以对每个数进行向右移动,移动的位数为,数组的下标,如果此时与1后仍大于0,证明这个位数为1,如果与1后为0则证明这位为0,不用选取
func subsets(nums []int) (sult [][]int) {
	length := len(nums)
	for nm := 0; nm < (1 << length); nm++ {
		var temp []int
		for i, v := range nums {
			if nm>>i&1 > 0 {
				temp = append(temp, v)
			}
		}
		sult = append(sult, temp)
	}
	return
}
