package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

// 弱智版本
func mergeInBetween2(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	l1 := list1
	for i := 0; i < a-1; i++ {
		l1 = l1.Next
	}

	l2 := list1
	for i := 0; i < b; i++ {
		l2 = l2.Next
	}

	p := list2
	for p.Next != nil {
		p = p.Next
	}

	l1.Next = list2
	p.Next = l2.Next

	return list1
}

// 优雅版本
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	l1 := list1
	var l2 *ListNode
	var i int

	for i != a-1 {
		i++
		l1 = l1.Next
	}

	l2 = l1.Next
	l1.Next = list2

	for i != b {
		i++
		l2 = l2.Next
	}

	for l1.Next != nil {
		l1 = l1.Next
	}

	l1.Next = l2

	return list1
}

func main() {
	l1 := &ListNode{Val: 0}
	x1 := &ListNode{Val: 1}
	x2 := &ListNode{Val: 2}
	x3 := &ListNode{Val: 3}
	x4 := &ListNode{Val: 4}
	x5 := &ListNode{Val: 5}
	x6 := &ListNode{Val: 6}
	l1.Next = x1
	x1.Next = x2
	x2.Next = x3
	x3.Next = x4
	x4.Next = x5
	x5.Next = x6

	l2 := &ListNode{
		Val: 10000,
		Next: &ListNode{
			Val: 10001,
			Next: &ListNode{
				Val: 10002,
				Next: &ListNode{
					Val:  10003,
					Next: nil,
				},
			},
		},
	}

	res := mergeInBetween(l1, 2, 5, l2)
	fmt.Println(res.String())

	//res2 := mergeInBetween2(l1, 2, 5, l2)
	//fmt.Println(res2.String())
}
