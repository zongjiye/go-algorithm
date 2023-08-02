package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 889 + 111 = 1000
// l1 = 9 -> 8 -> 8
// l2 = 1 -> 1 -> 1
// 1 + 999999

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 哨兵节点
	head := new(ListNode)
	l3 := head
	var temp int
	for l1 != nil || l2 != nil || temp != 0 {
		// 对应节点进行相加
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		// 结果取余， 满10进1
		l3.Next = &ListNode{
			Val: temp % 10,
		}
		l3 = l3.Next
		// 满10进1
		temp /= 10

	}

	return head.Next
}

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3.String())
}
