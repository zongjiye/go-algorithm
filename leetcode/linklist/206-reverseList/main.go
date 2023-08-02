package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

func reverseList(head *ListNode) *ListNode {
	pre := new(ListNode)
	p := head
	for p != nil {
		s := p
		p = p.Next
		s.Next = pre.Next
		pre.Next = s
	}

	return pre.Next
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3
	fmt.Println(n1.String())
	fmt.Println(reverseList(n1).String())
}
