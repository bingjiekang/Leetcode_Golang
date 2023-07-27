package main

import (
	"fmt"
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
	// var temp []rune = []rune{'0','1','2','3','4','5'}
	// fmt.Println(string(temp))
	fmt.Println(toHex(26), toHex(5), toHex(38), toHex(-1), toHex(-26), toHex(-18))
	// fmt.Println(string(temp[]))
	// fmt.Fprint()

}

func toHex(num int) string {
	var base []rune = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	var sult []rune = make([]rune, 0)
	temp := num
	if temp >= 0 {
		for temp/16 != 0 {
			sult = append(sult, base[temp%16])
			temp /= 16
		}
		sult = append(sult, base[temp])
		length := len(sult)
		for i, j := 0, length-1; i <= j; {
			sult[i], sult[j] = sult[j], sult[i]
			i++
			j--
		}
	} else {
		sult = []rune{'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f'}
		temp = (-1) * temp
		var tmi int = len(sult) - 1
		for ; temp/16 != 0; tmi-- {
			sult[tmi] = base[15-(temp%16)]
			temp /= 16
		}
		sult[tmi] = base[15-temp]
	}
	return string(sult)

}
