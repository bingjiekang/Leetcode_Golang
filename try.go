package main

import (
	"fmt"
)

func main() {
	var tm, tmn *ListNode
	tm.Val = 1
	tmn.Val = 1
	tm.Next = tmn
	fmt.Println(isPalindrome(tm))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var temp []int = make([]int, 1)
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}
	length := len(temp)
	left, right := 0, length-1
	for left < right {
		if temp[left] != temp[right] {
			return false
		}
		left++
		right--
	}
	return true
}
