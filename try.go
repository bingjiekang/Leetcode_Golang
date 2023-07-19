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
func longestPalindrome(s string) int {
	var hash map[rune]int = make(map[rune]int, 0)
	var sult, temp int
	for _, v := range s {
		hash[v]++
	}
	for i := range hash {
		if hash[i]%2 == 0 {
			sult += hash[i]
		} else {
			sult += (hash[i] - 1)
			temp = 1
		}
	}
	return sult + temp

}
