package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

// [a,b)区间的原地逆转
func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	for cur != b {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a, b := head, head
	// 找到b的位置
	for i := 0; i < k; i++ {
		// 不足k个的情况
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	fmt.Println(a.String())
	a.Next = reverseKGroup(b, k)
	return newHead
}

func main() {
	head := &ListNode{Val: 1}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 4}
	n4 := &ListNode{Val: 5}
	n5 := &ListNode{Val: 6}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	res := reverseKGroup(head, 2)
	fmt.Println(res.String())
}
