package main

// 要求:给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

// 解析:删除指定节点,如果第一个也相同的话,把头结点指向下一个,如果为空的话直接返回
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var Thead *ListNode
	Thead = head
	// 如果头结点不为空且第一个值相同将头指针指向下一个
	for head != nil && Thead.Val == val {
		head = Thead.Next
		Thead = head
	}
	// 为空直接返回
	if head == nil {
		return head
	}
	// 下一个节点不为空
	for Thead.Next != nil {
		// 如果下一个元素和值相同,则将下一个删除
		if Thead.Next.Val == val {
			Thead.Next = Thead.Next.Next
		} else { // 如果不同,则可以将当前节点指向下一个
			Thead = Thead.Next
		}
	}
	return head

}
