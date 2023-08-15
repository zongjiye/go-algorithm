package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}
	leftNode, cur := dummyHead, dummyHead
	for i := 1; i <= left; i++ {
		leftNode = cur
		cur = cur.Next
	}

	head = cur
	var pre *ListNode
	for i := left; i <= right; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	leftNode.Next, head.Next = pre, cur
	return dummyHead.Next
}

func main() {
	head := &ListNode{Val: 1}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 4}
	n4 := &ListNode{Val: 5}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	res := reverseBetween(head, 2, 4)
	fmt.Println(res.String())
}
