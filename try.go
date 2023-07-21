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
	fmt.Println(15 / 16)

}

func toHex(num int) string {
	var st []rune = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	var sult []rune = make([]rune, 0)
	if num >= 0 {
		temp := num
		for temp/16 != 0 {
			sult = append(sult, st[temp%16])
			temp /= 16
		}
		sult = append(sult, st[temp])
	} else {
		sult = []rune{'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f'}
		temp := num * (-1)
		for i := 7; temp/16 != 0 && i >= 0; i-- {
			sult[i] = st[16-(temp%16)]
			temp /= 16
		}

	}

}
