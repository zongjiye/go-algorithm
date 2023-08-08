package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

func removeElements(head *ListNode, val int) *ListNode {
	sentinel := new(ListNode)
	sentinel.Next = head
	s := sentinel
	for p := head; p != nil; p = p.Next {
		if p.Val == val {
			s.Next = p.Next
			continue
		}
		s = p
	}

	return sentinel.Next
}

func main() {
	head := &ListNode{Val: 1}
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 1}
	n3 := &ListNode{Val: 1}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	res := removeElements(head, 1)
	fmt.Println(res)
}
