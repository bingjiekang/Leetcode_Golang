package main

// 要求: 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代遍历翻转，1->2->3->4->5变成1<-2<-3<-4<-5，在遍历链表时，将当前节点的 next 指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点，在更改引用之前，还需要存储后一个节点，方便进行迭代，最后返回新的头引用。
func reverseList(head *ListNode) *ListNode {
	// 如果是空链表直接返回
	if head == nil {
		return head
	}
	// Lhead指向第一个节点
	Lhead := head
	// Mhead和Rhead指向第二个节点（可为空）
	Mhead := head.Next
	Rhead := head.Next
	// 当中间的节点为空时，证明已经走到了链表的最后，可以结束了
	for Mhead != nil {
		// 让右节点指向中间节点的右边
		Rhead = Rhead.Next
		// 让中间节点指向前一个节点进行翻转
		Mhead.Next = Lhead
		// 让前一个节点走向中间节点
		Lhead = Mhead
		// 让中间节点走向后一个节点
		Mhead = Rhead
	}
	// 将最开始的头结点的下一个节点指向nil即可
	head.Next = nil
	// 返回翻转后的头结点
	return Lhead
}

// 递归调用 a->b->c->d->e  则d->next->next(即e的next)=d,表示将d->e翻转成e->d.此时d还指向e,将d->next=nil置为空即可,这样递归到头结点,就不会形成环了(a就不会指向b,而是指向nil)
// func reverseList(head *ListNode) *ListNode {
//     // 如果是空链表或者递归的head的next为nil返回（边界）
//     if head == nil || head.Next == nil{
//         return head
//     }
//     // 递归遍历到最后一个节点，此时随着reverseList(head.Next)的返回，Thead指向最后一个节点,
//     Thead := reverseList(head.Next)
//     // 回到上一个栈里，此时head为前一个节点（重要），即上一个节点的head.next（从倒数第二个节点开始）
//     // 将此节点的下一个节点指向自己
//     head.Next.Next = head
//     // 将自己的下一个节点指向空
//     head.Next = nil
//     // Thread一直指向最后一个节点没有变化，变化的是head的每次递归返回
//     return Thead

// }
