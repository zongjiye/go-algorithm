package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast.Next != nil {
		return slow.Next
	}

	return slow
}

func main() {
	head := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	head.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	fmt.Println(middleNode(head))
	n5.Next = n6
	fmt.Println(middleNode(head))
}
