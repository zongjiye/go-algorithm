package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// NewLinkList 带头节点的链表
func NewLinkList() *ListNode {
	// 初始节点用于统计链表元素
	return &ListNode{Val: 0}
}

// NewListNode 节点
func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d", l.Val)
}

func (l *ListNode) Len() int {
	return l.Val
}

func (l *ListNode) Travel() {
	if l == nil {
		panic("link list not init")
	}
	// travel all nodes
	for p := l.Next; p != nil; p = p.Next {
		fmt.Println(p.String())
	}
}

func (l *ListNode) InitByHead(list []int) {
	if l == nil {
		panic("link list not init")
	}

	for _, v := range list {
		s := NewListNode(v)
		s.Next = l.Next
		l.Next = s
		l.Val++
	}
}

func (l *ListNode) InitByTail(list []int) {
	if l == nil {
		panic("link list not init")
	}

	p := l
	for _, v := range list {
		s := NewListNode(v)
		p.Next = s
		p = p.Next
		l.Val++
	}
}

func (l *ListNode) Insert(pos int, e int) {
	if l == nil {
		panic("link list not init")
	}
	if pos-1 < 0 || pos-1 > l.Len() {
		panic(fmt.Sprintf("index out of range [%d] with length %d", pos, l.Val))
	}

	p := l
	s := NewListNode(e)
	for i := 0; i < pos-1; i++ {
		p = p.Next
	}
	// insert
	s.Next = p.Next
	p.Next = s
	l.Val++
}

func (l *ListNode) Delete(pos int) int {
	if l == nil {
		panic("link list not init")
	}
	if pos-1 < 0 || pos > l.Len() {
		panic(fmt.Sprintf("index out of range [%d] with length %d", pos, l.Val))
	}

	p := l
	for i := 0; i < pos-1; i++ {
		p = p.Next
	}
	// delete
	v := p.Next.Val
	p.Next = p.Next.Next
	l.Val--
	return v
}

func main() {
	// create by head
	fmt.Println("create by head: ")
	head := NewLinkList()
	head.InitByHead([]int{1, 2, 3})
	head.Travel()
	// insert
	head.Insert(4, 99)
	fmt.Println("after insert: ")
	head.Travel()
	// delete
	v := head.Delete(4)
	fmt.Printf("after delete %d: \n", v)
	head.Travel()
	// create by tail
	fmt.Println("create by tail:")
	tail := NewLinkList()
	tail.InitByTail([]int{1, 2, 3})
	tail.Travel()
}
