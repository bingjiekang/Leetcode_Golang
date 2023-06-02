package main

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false

// 解析:将链表映射到列表里,对数据进行比较即可
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var temp []int
	// 对链表进行列表化输入
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}
	length := len(temp)
	left, right := 0, length-1
	// 如果不是回文返回false,否则返回true
	for left < right {
		if temp[left] != temp[right] {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
