package main

/*
要求:
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路:快慢指针,若是链表中没有环状链表,则慢指针永远追不上快指针,如果链表中存在环,则慢指针肯定可以追上快指针,快指针:fast = fast.Next.Next 慢指针:slow = slow.Next

func hasCycle(head *ListNode) bool {
	slow := head
	if head == nil {
		return false
	}
	// fast比slow快一个,方便后续继续便利和比较
	fast := head.Next
	for {
		// 如果出现空则证明到结尾,不存在环
		if fast == nil || slow == nil {
			return false
		} else if slow == fast { // 如果慢指针追上快指针则证明有环
			return true
		}
		slow = slow.Next
		if fast.Next == nil { // 防止.next.next越界,如果.next为空也直接返回
			return false
		}
		fast = fast.Next.Next
	}
}
