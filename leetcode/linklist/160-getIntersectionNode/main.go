package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 阁下掏出一个map 你该如何应对呢
func getIntersectionNodeByStupid(headA, headB *ListNode) *ListNode {
	mp := make(map[*ListNode]struct{})
	for p := headA; p != nil; p = p.Next {
		mp[p] = struct{}{}
	}

	for p := headB; p != nil; p = p.Next {
		if _, ok := mp[p]; ok {
			return p
		}
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}

	return l1
}

func main() {
	headA := &ListNode{Val: 1}
	headB := &ListNode{Val: 3}
	n1 := &ListNode{Val: 5}
	n2 := &ListNode{Val: 9}
	headA.Next = n1
	n1.Next = n2
	headB.Next = n2

	res := getIntersectionNodeByStupid(headA, headB)
	fmt.Printf("%v\n", res)

	res2 := getIntersectionNode(headA, headB)
	fmt.Printf("%v\n", res2)
}
