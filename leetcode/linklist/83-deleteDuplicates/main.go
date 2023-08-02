package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

// 弱智方法 我写了好久 废了
//	func deleteDuplicates(head *ListNode) *ListNode {
//		if head == nil {
//			return head
//		}
//		s := &ListNode{
//			Val: head.Val,
//		}
//		res := s
//		numMap := map[int]struct{}{
//			head.Val: {},
//		}
//
//		for p := head; p != nil; p = p.Next {
//			if _, ok := numMap[p.Val]; !ok {
//				s.Next = &ListNode{
//					Val: p.Val,
//				}
//				s = s.Next
//			}
//			numMap[p.Val] = struct{}{}
//		}
//
//		return res
//	}

// 此方法的奥妙之处真的除了图解我无法表达
func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil {
		s := curr
		for s != nil && s.Val == curr.Val {
			s = s.Next
		}
		curr.Next = s
		curr = curr.Next
	}

	return head
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	res := deleteDuplicates(l)
	fmt.Println(res.String())
}
