package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

// 纯弱智做法
// left 记录左半区数据
// right 记录右半区数据
func partition(head *ListNode, x int) *ListNode {
	right := new(ListNode)
	left := new(ListNode)
	l, r := left, right
	for p := head; p != nil; p = p.Next {
		if p.Val >= x {
			r.Next = &ListNode{Val: p.Val}
			r = r.Next
		} else {
			l.Next = &ListNode{Val: p.Val}
			l = l.Next
		}
	}
	l.Next = right.Next
	return left.Next
}

func main() {
	head := &ListNode{Val: 1}
	n1 := &ListNode{Val: 4}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 2}
	n4 := &ListNode{Val: 5}
	n5 := &ListNode{Val: 2}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	res := partition(head, 3)
	fmt.Println(res.String())
}
