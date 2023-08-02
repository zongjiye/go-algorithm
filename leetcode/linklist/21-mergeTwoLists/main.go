package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sentinel := new(ListNode)
	res := sentinel
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			sentinel.Next = list1
			list1 = list1.Next
		} else {
			sentinel.Next = list2
			list2 = list2.Next
		}

		sentinel = sentinel.Next
	}

	if list1 != nil {
		sentinel.Next = list1
	}

	if list2 != nil {
		sentinel.Next = list2
	}

	return res.Next
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l3 := mergeTwoLists(l1, l2)
	fmt.Println(l3.String())
}
