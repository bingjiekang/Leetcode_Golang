package main

import (
	"fmt"
	// "strconv"
	"math/bits"
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
	fmt.Println(bits.OnesCount(5))
	fmt.Fprint()

}

func readBinaryWatch(turnedOn int) (sult []string) {
	var i, j uint
	for i = 0; i < 12; i++ {
		for j = 0; j < 60; j++ {
			if bits.OnesCount(i)+bits.OnesCount(j) == turnedOn {
				sult = append(sult, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return
}
