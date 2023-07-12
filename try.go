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

//	func firstUniqChar(s string) int {
//		// 定义两个hash表，一个对字母的每一次下标进行赋值，一个记录第一次出现的字母下标,如果两个相同,证明没有重复,否则为重复
//		var list1, list2 map[rune]int
//		list1 = make(map[rune]int, 0)
//		list2 = make(map[rune]int, 0)
//		for i, v := range s {
//			list1[v] = i
//			_, err := list2[v]
//			if !err {
//				list2[v] = i
//			}
//		}
//		for i, v := range s {
//			if list1[v] == list2[v] {
//				return i
//			}
//		}
//		return -1
//	}
func firstUniqChar(s string) int {
	// 定义两个hash表，一个对字母的每一次下标进行赋值，一个记录第一次出现的字母下标,如果两个相同,证明没有重复,否则为重复
	var lt map[rune]int
	length := len(s)
	lt = make(map[rune]int, length)
	for i, v := range s {
		lt[v] = i
	}
	for i, v := range s {
		if lt[v] == i && lt[v] != -1 {
			return i
		}
		lt[v] = -1
	}
	return -1
}
