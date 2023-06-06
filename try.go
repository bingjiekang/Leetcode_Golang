package main

import (
	"fmt"
	"strings"
)

func main() {
	// var tm, tmn *ListNode
	// tm.Val = 1
	// tmn.Val = 1
	// tm.Next = tmn
	fmt.Println(wordPattern("a", "dog"))
}

func wordPattern(pattern string, s string) bool {
	var tmp []string = strings.Split(s, " ")
	var left map[rune]string = make(map[rune]string, 1)
	var right map[string]rune = make(map[string]rune, 1)
	if len(pattern) != len(tmp) {
		return false
	}
	for i, v := range pattern {
		_, ok := left[v]
		if !ok {
			left[v] = tmp[i]
			right[tmp[i]] = v
		} else {
			if left[v] != tmp[i] || right[tmp[i]] != v {
				return false
			} else if right[left[v]] != v || left[right[tmp[i]]] != tmp[i] {
				return false
			}
		}
	}
	return true
}
