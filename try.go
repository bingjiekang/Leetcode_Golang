package main

import (
	"fmt"
	"strings"
	// "strconv"
	// "time"
	// "math"
)

func main() {
	// // fmt.Println(countAndSay(5))
	// a := 3
	// tm := fmt.Sprintf("%b", a)
	// fmt.Println(time.Now().UnixNano())

	fmt.Println(reverseVowels("hello"))

}

func reverseVowels(s string) string {
	tmp := []rune(s)
	length := len(tmp)
	left, right := 0, length-1
	for left < right {
		for right > left && !strings.Contains("aeiouAEIOU", string(tmp[right])) {
			right--
		}
		for left < right && !strings.Contains("aeiouAEIOU", string(tmp[left])) {
			left++
		}
		if left < right {
			tmp[left], tmp[right] = tmp[right], tmp[left]
			left++
			right--
		} else {
			break
		}
	}
	return string(tmp)

}
