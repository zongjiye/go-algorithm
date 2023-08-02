package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

// 1 -> 2 -> 3 -> 4 -> 5, len = 5

// n = 1, del = 5
// n = 2, del = 4
// n = 3, del = 3
// n = 4, del = 2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// fast, slow
	fast := head
	slow := new(ListNode)
	slow.Next = head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 倒数第n个节点正好为第一个元素
	if fast == nil {
		return slow.Next.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	newHead := removeNthFromEnd(l1, 5)
	fmt.Println(newHead.String())
}
