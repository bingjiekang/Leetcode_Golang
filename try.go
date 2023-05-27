package main

func main() {
	// for Thead.Next != nil {
	// 	if Thead.Next.Val == val {
	// 		Thead.Next = Thead.Next.Next
	// 	}
	// 	Thead = Thead.Next
	// }
	// fmt.Println(AllAdd(19))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var Thead *ListNode
	Thead = head
	for Thead.Val == val {
		head = Thead.Next
		Thead = head
	}
	if head == nil {
		return head
	}
	for Thead.Next != nil {
		if Thead.Next.Val == val {
			Thead.Next = Thead.Next.Next
		}
		Thead = Thead.Next
	}
	return head

}
