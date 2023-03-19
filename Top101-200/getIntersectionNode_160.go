package main

/*
要求
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
判断两个链表是否相交，可以使用哈希集合存储链表节点.先将A链表加入哈希存储表，遍历B链表，如果B链表中的节点没有在哈希表中，证明这两个链表不相交，如果B链表中有节点在哈希表中，证明两个链表相交，那么返回这个相交节点即可
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 声明一个map用来实现哈希存储，key为节点，value为是否在里面
	HashHead := map[*ListNode]bool{}
	// 将A链表加入到哈希链表中
	for AHead := headA; AHead != nil; AHead = AHead.Next {
		HashHead[AHead] = true
	}
	// 遍历B链表，若B链表中有节点在哈希表中，证明两个链表相交
	for BHead := headB; BHead != nil; BHead = BHead.Next {
		if HashHead[BHead] {
			return BHead
		}
	}
	// 否则就不相交
	return nil
}
