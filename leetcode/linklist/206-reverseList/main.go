package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

// 头插法 原地逆转
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

// 快慢指针 原地逆转
// 在遍历链表时，将当前节点的 next 指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。
// 在更改引用之前，还需要存储后一个节点。最后返回新的头引用。
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3

	t1 := &ListNode{Val: 1}
	t2 := &ListNode{Val: 2}
	t3 := &ListNode{Val: 3}
	t1.Next = t2
	t2.Next = t3
	fmt.Println(reverseList(n1).String())
	fmt.Println(reverseList2(t1).String())
}
