package main

import (
	"fmt"
)

func main() {
	// var tm, tmn *ListNode
	// tm.Val = 1
	// tmn.Val = 1
	// tm.Next = tmn
	fmt.Println(isAnagram("anagram", "nagam"))
}

func isAnagram(s string, t string) bool {
	// var tmlf,tmrg map[byte]int,map[byte]int = make()
	var tmlf, tmrg map[rune]int = make(map[rune]int, 1), make(map[rune]int, 1)
	// var tmrg map[rune]int = make(map[rune]int, 1)
	for _, v := range s {
		_, ok := tmlf[v]
		if !ok {
			tmlf[v] = 1
		} else {
			tmlf[v]++
		}
	}
	for _, v := range t {

		_, ok := tmrg[v]
		if !ok {
			tmrg[v] = 1
		} else {
			tmrg[v]++
		}
	}
	for temp := range tmlf {
		if tmlf[temp] != tmrg[temp] {
			return false
		}
	}
	return true

}
