package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	for fast != nil && fast.Next != nil && head != nil {
		if fast == head {
			return true
		}
		fast, head = fast.Next.Next, head.Next
	}

	return false
}

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l2
	fmt.Println(hasCycle(l1))
}
